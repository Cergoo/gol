/*
	i18n pkg
	(c) 2014 Cergoo
	under terms of ISC license
*/

/*
 == 32
*/

package i18n

import (
	"fmt"
	"gol/err"
	"gol/filepath"
	"gol/jsonConfig"
	"gol/tplEngin/parser"
	"gol/tplEngin/plural"
	"io/ioutil"
	"strconv"
)

// types dictionarys define
type (
	Ti18n map[string]*tlang
	tlang struct {
		pluralRule plural.PluralRule       //
		plural     map[string][]string     // количествозависимое произношение
		phrase     map[string]*parser.Ttpl // строковые фразы
	}

	tTagText   []byte
	tTagVar    uint16
	tTagPlural struct {
		count uint16
		text  []string
	}

	TReplacer struct {
		langName string
		lang     *tlang
	}
)

// Create new replacer from language resources
func (t Ti18n) NewReplacer(langName string) *TReplacer {
	lang, e := t[langName]
	if !e {
		err.Panic(err.New("Not found lang resurse from langname: '"+langName+"'", 0))
	}
	return &TReplacer{langName: langName, lang: lang}
}

// Get lang
func (t *TReplacer) Lang() string {
	return t.langName
}

// Get phrase
func (t *TReplacer) P(key string, context ...interface{}) []byte {
	phrase, e := t.lang.phrase[key]
	if !e {
		return []byte(key)
	}
	if int(phrase.ContextLen) > len(context) {
		err.Panic(err.New("i18n Mismatch context len: lang:'"+t.Lang()+"', key:'"+key+"' ("+strconv.Itoa(int(phrase.ContextLen))+" , "+strconv.Itoa(len(context))+")", 0))
	}

	var result string
	for _, item := range phrase.Items {
		switch v := item.(type) {
		case tTagText:
			result += string(v)
		case tTagVar:
			result += fmt.Sprint(context[v])
		case *tTagPlural:
			result += v.text[t.lang.pluralRule(context[v.count].(float64))]
		}
	}
	return []byte(result)
}

// Get plural
func (t *TReplacer) Plural(key string, count float64) string {
	v, e := t.lang.plural[key]
	if e {
		return v[t.lang.pluralRule(count)]
	}
	return ""
}

// Create language resources
func Load(patch string) Ti18n {
	type (
		tmpLang struct {
			PluralRule string
			Plural     map[string][]string
			Phrase     map[string]string
		}
	)

	// создаётся временная структура и в неё парсится json
	tmpLangs := make(map[string]*tmpLang)
	fileList, e := ioutil.ReadDir(patch)
	err.Panic(e)

	var name string
	for _, item := range fileList {
		vtmpLang := new(tmpLang)
		vtmpLang.Plural = make(map[string][]string)
		jsonConfig.Load(patch+"/"+item.Name(), &vtmpLang)
		name, _ = filepath.Ext(item.Name())
		tmpLangs[name] = vtmpLang
	}

	i18n := make(Ti18n)
	toparse := new(parser.ToParse)
	toparse.Delimiter[0] = []byte("{{")
	toparse.Delimiter[1] = []byte("}}")
	toparse.ParseTag = parseTag
	toparse.ParseText = parseText

	for key, item := range tmpLangs {
		lang := new(tlang)
		lang.plural = item.Plural
		lang.pluralRule = plural.PluralRules[item.PluralRule]
		if lang.pluralRule == nil && len(lang.plural) > 0 {
			err.Panic(err.New("Not found plural rule: '"+item.PluralRule+"'", 0))
		}
		lang.phrase = make(map[string]*parser.Ttpl)
		for keyPhrase, itemPhrase := range item.Phrase {
			lang.phrase[keyPhrase] = parser.Parse([]byte(itemPhrase), toparse)
		}

		initAfterParse(lang, key)
		i18n[key] = lang
	}

	return i18n
}

func initAfterParse(lang *tlang, name string) {
	var (
		e   bool
		key string
	)
	// phrase loop
	for _, item := range lang.phrase {
		// tag loop
		for _, itemTag := range item.Items {
			switch v := itemTag.(type) {
			case *tTagPlural:
				key = v.text[0]
				v.text, e = lang.plural[key]
				if !e {
					err.Panic(err.New("Err parse:"+name+" Not found plural key: "+key, 0))
				}
			}
		}
	}
}

//  pars string to context id
func parseInt(source string) uint16 {
	i, e := strconv.Atoi(source)
	if e != nil || i < 0 {
		err.Panic(err.New("error parse to uint16: '"+source+"'", 0))
	}
	return uint16(i)
}

// parse plural tag
func parseTagPlural(source []string) (tag *tTagPlural, contextLen uint16) {
	if len(source) < 2 {
		err.Panic(err.New("error parsing to Plural Tag", 0))
	}
	tag = &tTagPlural{parseInt(source[1]), []string{source[0]}}
	contextLen = tag.count
	return
}

func parseText(source []byte) interface{} {
	return tTagText(source)
}

func parseTag(source []byte) (tag interface{}, contextLen uint16) {
	defer func() {
		if e := recover(); e != nil {
			v := e.(*err.OpenErr)
			v.Text += "err parse i18n: " + string(source)
			err.Panic(v)
		}
	}()

	list := parser.SplitWord(source, 32)
	switch list[0] {
	case "plural":
		tag, contextLen = parseTagPlural(list[1:])
	default:
		contextLen = parseInt(list[0])
		tag = tTagVar(contextLen)
	}
	return
}
