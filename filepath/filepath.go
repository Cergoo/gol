// (c) 2014 Cergoo
// under terms of ISC license

// Package filepath
package filepath

import (
	"os"
)

// Ext it's modified function standart "path/filepath" pkg
func Ext(fullname string) (name, ext string) {
	for i := len(fullname) - 1; i >= 0 && !os.IsPathSeparator(fullname[i]); i-- {
		if fullname[i] == '.' {
			return fullname[:i], fullname[i:]
		}
	}
	return "", ""
}
