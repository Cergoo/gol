package encodejson

import (
	"encoding/json"
	"fmt"
	"github.com/Cergoo/gol/test"
	//"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

type (
	tlickeByte byte

	t1 struct {
		F1 int
		F2 string
		f3 int // unexported field no encode no decode
		F4 []int
		F  *t2
		F5 int
		FN [4]byte
	}
	t2 struct {
		F1 string
		F  []int
		F2 string
	}
	tt struct {
		F1 string
		F2 string
	}
)

var (
	inInt        = int(17)
	inStr        = "tested string"
	inStr1       = ""
	inBoolt      = true
	inBoolf      = false
	inCopmlex128 = complex(float64(17.2), float64(112.1))
	inSlice      = []string{"1", "2", "3", "nnnn1", "nn2", "", "n1"}
	inSlice1     = []tlickeByte{1, 2, 3, 4}
	inArray      = [4]string{"1", "2", "3"}
	inArray1     = [4]uint8{1, 2, 3, 4}
	inStruct     = &t1{
		F1: -12,
		F2: "test1",
		f3: 100,
		F4: []int{1, 2, 4},
		F: &t2{
			F1: "test_str1рорн гшщзй☂",
			F2: "test_str2",
		},
		F5: 12,
		FN: [4]byte{1, 2, 3, 4},
	}
	inMap            = map[int]string{1: "f1", 2: "f2", 4: "f4"}
	inMap1           = map[int]*t2{1: &t2{"f1", nil, "f2"}, 2: &t2{"f2", nil, "f2"}, 3: nil}
	inMap2           = map[int]t2{1: t2{"f1", []int{}, "f2"}, 2: t2{"f2", []int{}, "f2"}}
	inMapN           = map[tt]int{tt{"f1", "f2"}: 1, tt{"f2", "f2"}: 2}
	inMapInterface   = map[int]interface{}{1: 1, 2: tt{"f2", "f2"}, 3: &t2{"f2", nil, "f2"}, 4: nil, 5: time.Now().UTC()}
	inSliceInterface = []interface{}{-1, 100, "nnnnnn", nil, 7.5, time.Now().UTC(), tt{"f2", "f2"}, &t2{"f2", nil, "f2"}, []byte{12, 17, 0, 18}, []int{12, 10, 17}}
)

func TestED(t *testing.T) {
	var (
		buf, buf1 []byte
	)
	t1 := test.New(t)

	buf = Encode(buf, inInt)
	buf1, _ = json.Marshal(inInt)
	t1.Eq(buf, buf1)
	buf = buf[:0]

	buf = Encode(buf, inSlice)
	buf1, _ = json.Marshal(inSlice)
	t1.Eq(buf, buf1)
	buf = buf[:0]

	buf = Encode(buf, inStruct)
	buf1, _ = json.Marshal(inStruct)
	t1.Eq(buf, buf1)
	fmt.Println("nn", string(buf))
	fmt.Println("nn", string(buf1))
	buf = buf[:0]
}
