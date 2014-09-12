package main

import (
	"fmt"
	"github.com/Cergoo/gol/binaryED/binaryED"
	"github.com/Cergoo/gol/fastbuf"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	filename := "dump"
	data := []interface{}{-1, 100, "str1", 7.5, time.Now().UTC(), nil, uint8(2), []int{12, 10, 17}}

	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	buf := fastbuf.New(nil, 0, fp)
	binaryED.Encode(buf, data)
	buf.FlushToWriter()
	fp.Close()

	data = data[:0]
	dump, _ := ioutil.ReadFile(filename)
	buf = fastbuf.New(dump, 0, nil)
	decoder := binaryED.NewDecoder(buf)
	decoder.Decode(&data)

	fmt.Println(data)
}
