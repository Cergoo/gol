/*
	filepath pkg
	(c) 2014 Cergoo
	under terms of ISC license
*/

package filepath

import (
	"os"
)

//	modified function Ext standart "path/filepath" pkg
func Ext(fullname string) (name, ext string) {
	for i := len(fullname) - 1; i >= 0 && !os.IsPathSeparator(fullname[i]); i-- {
		if fullname[i] == '.' {
			return fullname[:i], fullname[i:]
		}
	}
	return "", ""
}
