package cache

import (
	"io"
)

type (
	TCortege struct {
		Key string
		Val interface{}
	}

	Cache interface {
		GetBucketsStat() (countitem uint64, countbucket uint32, stat [][2]int)
		Get(string) interface{}
		Set(string, interface{}) bool
		Del(string) (val interface{})
		DelAll()
		Range(chan<- *TCortege)
		Inc(string, float64) interface{}
		Save(io.Writer) error
		SaveFile(string) error
		Load(io.Reader) error
		LoadFile(string) error
		Len() I_counter
	}

	I_counter interface {
		Get() uint64
		GetLimit() uint64
		SetLimit(v uint64)
		Check() bool
	}
)
