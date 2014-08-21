package cache

import (
	"io"
)

// Returned result code
const (
	ResultExist = iota
	ResultAdd
	ResultNoExistNoAdd
)

// Mode operation set
const (
	ModeSet_OnlyUpdate = iota
	ModeSet_OnlyInsert
	ModeSet_UpdateOrInsert
)

type (
	// (key, value) cortege
	TCortege struct {
		Key string
		Val interface{}
	}

	Cache interface {
		GetBucketsStat() (countitem uint64, countbucket uint32, stat [][2]int)
		Get(string) interface{}
		Set(cortege *TCortege, mode uint8) (val interface{}, actionResult uint8)
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

	// Counter interface
	I_counter interface {
		Get() uint64
		GetLimit() uint64
		SetLimit(v uint64)
		Check() bool
	}
)
