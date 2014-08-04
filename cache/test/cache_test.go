package cache

import (
	"errors"
	"github.com/Cergoo/gol/cache"
	"github.com/Cergoo/gol/hash"
	"strconv"
	"testing"
	"time"
)

var (
	cache1 cache.Cache
)

func init() {
	cache1 = cache.New(hash.HashFAQ6, 0, true, 5*time.Second, nil)
}

func Test_Get(t *testing.T) {
	var v int
	for i := 0; i < 100000; i++ {
		cache1.Set("item"+strconv.Itoa(i), i)
	}

	for i := 0; i < 100000; i++ {
		v = cache1.Get("item" + strconv.Itoa(i)).(int)
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
	cache1.Inc("item101", 10)
	v = cache1.Get("item101")
	if v != 111 {
		t.Error("err")
		return
	}
	cache1.Inc("item101", -11)
	v = cache1.Get("item101")
	if v != 100 {
		t.Error("err")
		return
	}
}

func Test_SetFunc(t *testing.T) {
	var (
		v interface{}
	)

	arg := 2
	f := func(val interface{}) (interface{}, error) {
		var (
			v int
			b bool
		)
		v, b = val.(int)
		if !b {
			return nil, errors.New("Mismatch type")
		}
		return v * arg, nil
	}

	cache1.SetFunc("item101", 1, f)
	v = cache1.Get("item101")
	if v != 200 {
		t.Error("err", v)
		return
	}

}

func Test_Del(t *testing.T) {
	cache1.Del("item1")
	v := cache1.Get("item1")
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
	cache1.Set("item10", 11)
	err = cache1.LoadFile("f")
	if err != nil {
		t.Error("err")
		return
	}
	v := cache1.Get("item10")
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
