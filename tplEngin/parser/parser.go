// (c) 2014 Cergoo
// under terms of ISC license

// Package parser it's a util from i18n & tpl pkg
package parser

import (
	"bytes"
)

type (

	/*
		ToParse struct describes a tool for parsing:
		    Delimiter - left and right tag gelimiter;
		    ParseText - function of parse text;
		    ParseTag  - function of parse tag.
	*/
	ToParse struct {
		Delimiter           [2][]byte
		ParseText, ParseTag func([]byte) interface{}
	}

	// Items of template
	Ttpl []interface{}
)

/*
	FindTag search of a tag in slice.
	Return: leftpart, tag value, end tag number, success.
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

// SplitWord split a []byte into words by delimiters, cut repeatable delimiters
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

// StrPrefix if (header(a) == b) { trim head and return tail } else { return nil }
func StrPrefix(a []byte, b string) []byte {
	if len(a) >= len(b) && string(a[:len(b)]) == b {
		return a[len(b):]
	}
	return nil
}

// Parse universal parse metode, return template
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
