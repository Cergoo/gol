package main

import (
	"fmt"
	"gol/tplEngin/i18n"
)

func main() {
	lang := i18n.New()
	lang.Load("lang", true)
	lang.Load("lang1", true)

	fmt.Println("ru")
	replacer := lang.NewReplacer("ru")
	fmt.Println(string(replacer.P("message")))
	fmt.Println(string(replacer.P("message1", "поле1", float64(2))))
	fmt.Println(2.5, replacer.Plural("apple", float64(2.5)))
	fmt.Println(string(replacer.P("pkgLang1_message10")))

	fmt.Println("en")
	replacer = lang.NewReplacer("en")
	fmt.Println(string(replacer.P("message")))
	fmt.Println(string(replacer.P("message1", "поле1", float64(2))))
	fmt.Println(2.5, replacer.Plural("apple", float64(2.5)))
	fmt.Println(string(replacer.P("pkgLang1_message10")))

}
