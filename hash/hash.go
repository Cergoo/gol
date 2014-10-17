// (c) 2014 Cergoo
// under terms of ISC license

// Package hash it's a non-cryptographic hash 32 functions library.
package hash

import (
	"unsafe"
)

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
)

// Murmur3 hash, get from github.com/spaolacci/murmur3
func Murmur3(data []byte) (h uint32) {
	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}
	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32
		h ^= k1
		h = (h << 13) | (h >> 19) // rotl32(h1, 13)
		h = h*5 + 0xe6546b64
	}

	tail := data[nblocks*4:]
	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32
		h ^= k1
	}

	h ^= uint32(len(data))
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return
}

// Jenkins hash
func HashJenkins(str []byte) (h uint32) {
	for i := range str {
		h += uint32(str[i])
		h += h << 10
		h ^= h >> 6
	}
	h += h << 3
	h ^= h >> 11
	h += h << 15
	return
}

// HashRot13
func HashRot13(str []byte) (h uint32) {
	for i := range str {
		h += uint32(str[i])
		h -= (h << 13) | (h >> 19)
	}
	return
}

// HashLy
func HashLy(str []byte) (h uint32) {
	for i := range str {
		h = (h * 1664525) + uint32(str[i]) + 1013904223
	}
	return
}

// HashRs
func HashRs(str []byte) (h uint32) {
	const b = 378551
	var a uint32 = 63689
	for i := range str {
		h = h*a + uint32(str[i])
		a *= b
	}
	return
}
