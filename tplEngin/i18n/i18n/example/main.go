// Example use pkg
package main

import (
	"fmt"
	"github.com/Cergoo/gol/tplEngin/i18n/i18n"
	"github.com/Cergoo/gol/tplEngin/i18n/i18n/mod/humanize"
)

type (
	item struct {
		name  string
		count float32
	}

	formatted float64
)

// fmt.Stringer implimentation
func (t formatted) String() string {
	return fmt.Sprintf("%20.5f", t)
}

func main() {
	lang := i18n.New()
	lang.Load("lang")
	lang.Load("lang1")
	humanize.Init(lang, "human/lang", humanize.FHumanByte|humanize.FHumanByteLong)
	lang.Init(true)

	names := []*item{
		&item{name: "UserName1", count: 12.2},
		&item{name: "UserName2", count: 12.25},
		&item{name: "UserName3", count: 101},
		&item{name: "UserName4", count: 120000},
	}

	replacerRuLang, _ := lang.NewReplacer("ru")
	replacerEnLang, _ := lang.NewReplacer("en")

	for _, v := range names {
		fmt.Println("ru: ", string(replacerRuLang.P("message2", v.name, v.count)))
		fmt.Println("en: ", string(replacerEnLang.P("message2", v.name, v.count)))
		fmt.Println("ru: ", string(replacerRuLang.P("message3", v.name, v.count)))
		fmt.Println("en: ", string(replacerEnLang.P("message3", v.name, v.count)))
		fmt.Println("ru: ", string(replacerRuLang.P("message4", v.name, v.count)))
		fmt.Println("en: ", string(replacerEnLang.P("message4", v.name, v.count)))
	}

	fmt.Println(formatted(2.5), replacerRuLang.Plural("apple", 2.5))
	fmt.Println(2.5, replacerEnLang.Plural("apple", 2.5))

}
