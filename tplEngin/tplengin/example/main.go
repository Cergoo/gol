package main

import (
	"fmt"
	"gol/fastbuf"
	"gol/tplEngin/tplengin"
)

func main() {
	tple := tplengin.New()
	tple.I18n.Load("tpls/maintpls/lang", false)
	tple.Load("tpls/maintpls/maintpls")
	tple.InitBeforeUse()
	r := tple.NewReplacer("en")

	context := make(map[string]interface{})
	context["a"] = 10
	context["f1"] = []interface{}{float64(12)}

	buf := &fastbuf.Buf{}
	r.Replace("/maintpls/1.tpl", context, buf)
	fmt.Println(string(buf.Flush()))
}
