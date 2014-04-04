package test

import (
	"fmt"
	"reflect"
)

type TT bool

func (t *TT) Eq(id string, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		fmt.Println("err:", id, a, b)
		*t = true
	}
}
