package main

import (
	"fmt"
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
	context["f1"] = []interface{}{"MyText", float64(12)}

	rezult := make([]byte, 0)
	rezult = r.Replace("/maintpls/1.tpl", context, rezult)
	fmt.Println(string(rezult))
}
