/*
  support comments in json
  (c)	2013 Cergoo
	under terms of ISC license
*/

package jsonConfig

import (
	"encoding/json"
	"gol/err"
	"io/ioutil"
	"regexp"
)

func Load(fromPath string, toVar interface{}) {
	file, e := ioutil.ReadFile(fromPath)
	err.Panic(e)
	reg := regexp.MustCompile(`(?im)(//[^"'}]*$)|(?s)(/\*[^"']*?\*/\s*$)`)
	fileRezult := reg.ReplaceAll(file, []byte{})
	err.Panic(json.Unmarshal(fileRezult, toVar))
}
