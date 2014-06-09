/*
	parser util from i18n & tpl pkg
	(c) 2014 Cergoo
	under terms of ISC license
*/

package parser

import (
	"bytes"
	"gol/err"
	"strconv"
)

type (

	/*
		структура описывающая инструмент для парсинга
		Delimiter - left and right tag gelimiter
		ParseText - function of parse text
		ParseTag  - function of parse tag
	*/
	ToParse struct {
		Delimiter           [2][]byte
		ParseText, ParseTag func([]byte) interface{}
	}

	Ttpl []interface{} // items of template
)

/*
	search of tag in a slice
	return of leftpart, tag value, end tag number, success
*/
func FindTag(source []byte, delimiter [2][]byte) (lpart, tag []byte, end int, success bool) {
	var begin int
	// look a delimiter beginning
	begin = bytes.Index(source, delimiter[0])
	if begin > -1 {
		lpart = source[:begin]
		// look a delimiter ending
		begin = begin + len(delimiter[0])
		source = source[begin:]
		end = bytes.Index(source, delimiter[1])
		if end > -1 {
			tag = bytes.TrimSpace(source[:end])
			end += len(delimiter[1]) + begin
			success = true
		}
	}
	return
}

/*
	split a []byte into words by delimiters
*/
func SplitWord(source []byte, delimiters byte) []string {
	var (
		begin int
	)
	result := make([]string, 0)
	find := false
	for i := range source {
		if (source[i] < 33 && delimiters == 32) || source[i] == delimiters {
			if find {
				result = append(result, string(source[begin:i]))
				find = false
			}
		} else {
			if !find {
				begin = i
				find = true
			}
		}
	}

	if find {
		result = append(result, string(source[begin:]))
	}
	return result
}

/*
	if header(a) == b trim head and return tail
	else return nil
*/
func StrPrefix(a []byte, b string) []byte {
	if len(a) >= len(b) && string(a[:len(b)]) == b {
		return a[len(b):]
	}
	return nil
}

/*
	Universal parse metode, return template
*/
func Parse(source []byte, toparse *ToParse) (tpl Ttpl) {
	var (
		success     bool
		lpart, tag  []byte
		end, newend int
		ptag        interface{}
	)

	lpart, tag, end, success = FindTag(source, toparse.Delimiter)
	for success {
		if len(lpart) > 0 {
			tpl = append(tpl, toparse.ParseText(lpart))
		}
		ptag = toparse.ParseTag(tag)
		if ptag != nil {
			tpl = append(tpl, ptag)
		}

		lpart, tag, newend, success = FindTag(source[end:], toparse.Delimiter)
		end += newend
	}
	lpart = source[end:]
	if len(lpart) > 0 {
		tpl = append(tpl, toparse.ParseText(lpart))
	}

	return
}

//  pars string to context id
func ParseInt(source string) uint {
	i, e := strconv.Atoi(source)
	if e != nil || i < 0 {
		err.Panic(err.New("error parse to uint: '"+source+"'", 0))
	}
	return uint(i)
}
