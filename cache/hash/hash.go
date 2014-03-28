package hash

// partitioner
func Keytoid(key []byte, l int) uint32 {
	// h & (len - 1)
	return hash(key) % uint32(l)
	//return hash(key) & (uint32(l) - 1)
}

func hash(str []byte) uint32 {
	var (
		h uint32
	)
	for i := 0; i < len(str); i++ {
		h += uint32(str[i])
		h += h << 10
		h ^= h >> 6
	}
	//h += h << 3
	//h ^= h >> 11
	//h += h << 15
	return h
}
