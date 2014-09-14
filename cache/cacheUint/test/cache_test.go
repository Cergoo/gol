package cacheUint

import (
	"github.com/Cergoo/gol/cache/cacheUint"
	"testing"
	"time"
)

var (
	cache1 cacheUint.Cache
)

func init() {
	cache1 = cacheUint.New(true, 20*time.Second, nil)
}

func Test_Get(t *testing.T) {
	var v int
	for i := 0; i < 100000; i++ {
		cache1.Set(uint64(i), i, 1, cacheUint.UpdateOrInsert)
	}

	for i := 0; i < 100000; i++ {
		v = cache1.Get(uint64(i)).(int)
		if v != i {
			t.Error("err")
			return
		}
	}

}

func Test_Inc(t *testing.T) {
	var (
		v interface{}
	)
	cache1.Inc(uint64(101), 10)
	v = cache1.Get(uint64(101))
	if v != 111 {
		t.Error("err")
		return
	}
	cache1.Inc(uint64(101), -11)
	v = cache1.Get(uint64(101))
	if v != 100 {
		t.Error("err")
		return
	}
}

func Test_Del(t *testing.T) {
	cache1.Del(uint64(1))
	v := cache1.Get(uint64(1))
	if v != nil {
		t.Error("err")
		return
	}
}

func Test_SaveLoad(t *testing.T) {
	var (
		err error
	)
	err = cache1.SaveFile("f")
	if err != nil {
		t.Error("err")
		return
	}
	cache1.Set(uint64(10), 11, 1, cacheUint.UpdateOrInsert)
	err = cache1.LoadFile("f")
	if err != nil {
		t.Error("err")
		return
	}
	v := cache1.Get(uint64(10))
	if v != 10 {
		t.Error("err", v)
		return
	}
}

func Test_Len(t *testing.T) {
	if cache1.Len().Get() != 99999 {
		t.Error("err")
		return
	}
}

func Test_janitar(t *testing.T) {
	time.Sleep(12 * time.Second)
	if cache1.Len().Get() == 99999 {
		t.Error("err")
		return
	}
}
