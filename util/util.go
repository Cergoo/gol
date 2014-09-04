// (c) 2014 Cergoo
// under terms of ISC license

// Package util it's go util pkg
package util

// Min get min int
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max get max int
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// MaxSet set max to var
func MaxSet(a *int, b int) {
	if *a < b {
		*a = b
	}
}

// MaxSet set min to var
func MinSet(a *int, b int) {
	if *a > b {
		*a = b
	}
}
