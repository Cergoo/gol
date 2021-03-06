/*
Package fastutil it's a util from fast geneate encoder/decoder
*/
package fastutil

import (
	"strings"
)

// TypeName get name pkg variable
func TypeName(val string, pointered bool) string {
	if !pointered && val[0] == '*' {
		val = val[1:]
	}
	val = strings.Replace(val, "]main.", "]", -1)
	val = strings.Replace(val, "*main.", "*", -1)
	val = strings.Replace(val, "[main.", "[", -1)
	parts := strings.SplitN(val, ".", 2)
	if len(parts) > 1 {
		if parts[0] == "main" || parts[0] == "*main" {
			if parts[0][0] == '*' {
				return "*" + parts[1]
			}
			return parts[1]
		}
	}
	return val
}
