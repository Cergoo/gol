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

// PathEndSeparator check a PathSeparator into path end
func PathEndSeparator(path string) string {
	if path[len(path)-1] == os.PathSeparator {
		return path
	}
	return path + string(os.PathSeparator)
}
