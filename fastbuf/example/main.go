// Example use pkg.
package main

import (
	"fmt"
	"github.com/Cergoo/gol/fastbuf"
)

func main() {
	b := fastbuf.New(nil, 0, nil)
	b.Write([]byte("12"))
	fmt.Print(b.FlushP())

}
