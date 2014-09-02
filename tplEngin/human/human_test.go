package human

import (
	"github.com/Cergoo/gol/test"
	"testing"
)

type testData struct {
	Size   uint64
	Result string
}

var tests = []testData{
	{4, "4B"},
	{10, "10B"},
	{909, "909B"},
	{1024, "1.0KB"},
	{206848, "202KB"},
	{1048576, "1.0MB"},
	{1439744, "1.4MB"},
	{353974272, "338MB"},
	{1073741824, "1.0GB"},
	{126419740672, "118GB"},
	{1099511627776, "1.0TB"},
	{10239999998976, "9.3TB"},
	{942079897600000, "857TB"},
	{1125899906842624, "1.0PB"},
	{93429759897600000, "83PB"},
	{1152921504606846976, "1.0EB"},
	{6314666666666665984, "5.5EB"},
}

func TestSize(t *testing.T) {
	t1 := test.New(t)
	for _, pair := range tests {
		t1.Eq(Byten(pair.Size), pair.Result)
	}
}
