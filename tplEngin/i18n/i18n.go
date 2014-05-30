/*
	i18n pkg
	(c) 2014 Cergoo
	under terms of ISC license
*/

package i18n

import (
	"fmt"
	"gol/err"
	gfilepath "gol/filepath"
	"gol/jsonConfig"
	"gol/refl"
	"gol/tplEngin/i18n/plural"
	"gol/tplEngin/parser"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

// types dictionarys define
type (
	Ti18n map[string]*tlang
	tlang struct {
		pluralRule plural.PluralRule   // language plural rule
		plural     map[string][]string // plural pronunciation
		items      map[string]*titem   // tpl
	}
	titem struct {
		items        parser.Ttpl // tags
		contextCount int         // context var count
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

// Print phrase from map store.
func (t *TReplacer) P(key string, context ...interface{}) []byte {
	tpl := t.lang.items[key]
	if tpl == nil {
		return nil
	}
	return t.p(tpl, context)
}

// Get plural. Use if Load (pluralAccess)
func (t *TReplacer) Plural(key string, count float64) string {
	v, e := t.lang.plural[key]
	if e {
		return v[t.lang.pluralRule(count)]
	}
	return ""
}

// Create language obj
func New() Ti18n {
	return make(Ti18n)
}

// Loade language resources
func (t Ti18n) Load(patch string, pluralAccess bool) {
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

	var (
		name   string
		valPre map[string]string
		keyPre string
	)

	for _, item := range fileList {
		vtmpLang := new(tmpLang)
		vtmpLang.Plural = make(map[string][]string)
		jsonConfig.Load(patch+string(filepath.Separator)+item.Name(), &vtmpLang)
		name, _ = gfilepath.Ext(item.Name())
		tmpLangs[name] = vtmpLang
	}

	// chek equivalent all lang resurce
	for key, val := range tmpLangs {
		if valPre != nil && !refl.MapKeysEq(valPre, val.Phrase) {
			err.Panic(err.New("Lang phrase not equivalent: "+keyPre+", "+key, 0))
		}
		valPre = val.Phrase
		keyPre = key
	}

	toparse := new(parser.ToParse)
	toparse.Delimiter[0] = []byte("{{")
	toparse.Delimiter[1] = []byte("}}")
	toparse.ParseTag = parseTag
	toparse.ParseText = parseText

	for key, item := range tmpLangs {
		lang := new(tlang)
		lang.items = make(map[string]*titem)
		lang.plural = item.Plural
		lang.pluralRule = plural.PluralRules[item.PluralRule]
		if lang.pluralRule == nil && len(lang.plural) > 0 {
			err.Panic(err.New("Not found plural rule: '"+item.PluralRule+"'", 0))
		}

		for keyPhrase, itemPhrase := range item.Phrase {
			lang.items[keyPhrase] = &titem{items: parser.Parse([]byte(itemPhrase), toparse), contextCount: -1}
		}

		initAfterParse(lang, key)
		if !pluralAccess {
			lang.plural = nil
		}

		existLang := t[key]
		if existLang == nil {
			t[key] = lang
		} else {
			// add phrase
			for key, val := range lang.items {
				existLang.items[key] = val
			}
			// add plural
			for key, val := range lang.plural {
				existLang.plural[key] = val
			}
		}
	}
}

// Get phrase
func (t *TReplacer) p(tpl *titem, context []interface{}) []byte {
	if int(tpl.contextCount) > len(context) {
		err.Panic(err.New("i18n Mismatch context len: ("+strconv.Itoa(int(tpl.contextCount))+" , "+strconv.Itoa(len(context))+")", 0))
	}

	var result []byte
	for _, item := range tpl.items {
		switch v := item.(type) {
		case tTagText:
			result = append(result, v...)
		case tTagVar:
			result = append(result, []byte(fmt.Sprint(context[v]))...)
		case *tTagPlural:
			result = append(result, []byte(v.text[t.lang.pluralRule(context[v.count].(float64))])...)
		}
	}
	return result
}

// Init plural tag
func initAfterParse(lang *tlang, name string) {
	var (
		e   bool
		key string
	)
	// phrase loop
	for _, item := range lang.items {
		// tag loop
		for _, itemTag := range item.items {
			switch v := itemTag.(type) {
			case *tTagPlural:
				key = v.text[0]
				if int(v.count) > item.contextCount {
					item.contextCount = int(v.count)
				}
				v.text, e = lang.plural[key]
				if !e {
					err.Panic(err.New("Err parse:"+name+" Not found plural key: "+key, 0))
				}
			case tTagVar:
				if int(v) > item.contextCount {
					item.contextCount = int(v)
				}
			}
		}
		item.contextCount++
	}
}

// parse plural tag
func parseTagPlural(source []string) *tTagPlural {
	if len(source) < 2 {
		err.Panic(err.New("error parsing to Plural Tag: "+fmt.Sprint(source), 0))
	}
	return &tTagPlural{parser.ParseInt(source[1]), []string{source[0]}}
}

func parseText(source []byte) interface{} {
	return tTagText(source)
}

func parseTag(source []byte) interface{} {
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
		return parseTagPlural(list[1:])
	default:
		return tTagVar(parser.ParseInt(list[0]))
	}
}
