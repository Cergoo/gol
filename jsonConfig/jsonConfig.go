/*
  support comments in json
  (c)	2013-2014 Cergoo
  under terms of ISC license
*/

/*
" == 34
/ == 47
\ == 92
* == 42
*/

package jsonConfig

import (
	"encoding/json"
	"gol/err"
	"io/ioutil"
)

// Load & remove comments from source .json file
func Load(fromPath string, toVar interface{}) {
	file, e := ioutil.ReadFile(fromPath)
	err.Panic(e)
	file = RemoveComment(file)
	err.Panic(json.Unmarshal(file, toVar))
}

// Remove comments from source .json
func RemoveComment(source []byte) (result []byte) {
	var (
		stateBlok, stateComment1, stateComment2 bool
	)

	for i := 0; i <= len(source)-1; i++ {

		if stateBlok {
			result = append(result, source[i])
			if beginOrEndBlok(source, i) {
				stateBlok = false
			}
			continue
		}

		if stateComment1 {
			if endComment1(source, i) {
				stateComment1 = false
			}
			continue
		}

		if stateComment2 {
			if endComment2(source, i) {
				stateComment2 = false
				i++
			}
			continue
		}

		if beginOrEndBlok(source, i) {
			result = append(result, source[i])
			stateBlok = true
			continue
		}

		if beginComment1(source, i) {
			stateComment1 = true
			i++
			continue
		}

		if beginComment2(source, i) {
			stateComment2 = true
			i++
			continue
		}

		result = append(result, source[i])

	}
	return
}

// detect "
func beginOrEndBlok(source []byte, i int) bool {
	return source[i] == 34 && (i == 0 || source[i-1] != 92)
}

// detect //
func beginComment1(source []byte, i int) bool {
	return i < len(source) && source[i] == 47 && source[i+1] == 47
}

// detect /n
func endComment1(source []byte, i int) bool {
	return source[i] == 10
}

// detect /*
func beginComment2(source []byte, i int) bool {
	return i < len(source) && source[i] == 47 && source[i+1] == 42
}

// detect */
func endComment2(source []byte, i int) bool {
	return i < len(source) && source[i] == 42 && source[i+1] == 47
}
