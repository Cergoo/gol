/*
	go util pkg
	(c) 2014 Cergoo
	under terms of ISC license
*/
package util

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
