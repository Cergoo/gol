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
	"gol/refl"
	"gol/tplEngin/i18n/plural"
	"gol/tplEngin/parser"
	"io/ioutil"
	"strconv"
	"strings"
)

// types dictionarys define
type (
	Ti18n map[string]*tlang
	tlang struct {
		pluralRule  plural.PluralRule       // language plural rule
		plural      map[string][]string     // plural pronunciation
		phraseMap   map[string]*parser.Ttpl // phrases map store
		phraseSlice []*parser.Ttpl          // phrases slice store
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

// Print. Get phrase from map store. Use if Load (mapAccess)
func (t *TReplacer) P(key string, context ...interface{}) []byte {
	tpl, e := t.lang.phraseMap[key]
	if e {
		return t.p(tpl, context, key)
	}
	return nil
}

// Print faste. Get phrase from slice store. Use if Load (sliceAccess)
func (t *TReplacer) Pf(key int, context ...interface{}) []byte {
	if len(t.lang.phraseSlice) > key {
		return t.p(t.lang.phraseSlice[key], context, strconv.Itoa(key))
	}
	return nil

}

// Get plural. Use if Load (pluralAccess)
func (t *TReplacer) Plural(key string, count float64) string {
	v, e := t.lang.plural[key]
	if e {
		return v[t.lang.pluralRule(count)]
	}
	return ""
}

// Create language resources
func Load(patch string, pluralAccess bool) Ti18n {
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
		parts  []string
		tpl    *parser.Ttpl
		id     int
		valPre map[string]string
		keyPre string
	)

	for _, item := range fileList {
		vtmpLang := new(tmpLang)
		vtmpLang.Plural = make(map[string][]string)
		jsonConfig.Load(patch+"/"+item.Name(), &vtmpLang)
		name, _ = filepath.Ext(item.Name())
		tmpLangs[name] = vtmpLang
	}

	// chek equivalent all lang resurce
	for key, val := range tmpLangs {
		if valPre != nil && !refl.MapKeysEq(valPre, val.Phrase) {
			err.Panic(err.New("lang prase not equivalent: "+keyPre+", "+key, 0))
		}
		valPre = val.Phrase
		keyPre = key
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
		lang.phraseMap = make(map[string]*parser.Ttpl)
		lang.phraseSlice = make([]*parser.Ttpl, len(item.Phrase))
		for keyPhrase, itemPhrase := range item.Phrase {
			tpl = parser.Parse([]byte(itemPhrase), toparse)
			keyPhrase = strings.TrimSpace(keyPhrase)
			parts = strings.SplitN(keyPhrase, " ", 2)
			id, e = strconv.Atoi(parts[0])
			if e == nil {
				tpl.Id = uint16(id)
				lang.phraseSlice[id] = tpl
			}

			if len(parts) > 1 {
				lang.phraseMap[strings.TrimSpace(parts[1])] = tpl
			} else {
				lang.phraseMap[keyPhrase] = tpl
			}
		}

		initAfterParse(lang, key)
		if !pluralAccess {
			lang.plural = nil
		}
		i18n[key] = lang
	}

	return i18n
}

// Get phrase
func (t *TReplacer) p(tpl *parser.Ttpl, context []interface{}, key string) []byte {
	if int(tpl.ContextLen) > len(context) {
		err.Panic(err.New("i18n Mismatch context len: lang:'"+t.Lang()+"', key:'"+key+"' ("+strconv.Itoa(int(tpl.ContextLen))+" , "+strconv.Itoa(len(context))+")", 0))
	}

	var result string
	for _, item := range tpl.Items {
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

func initAfterParse(lang *tlang, name string) {
	var (
		e   bool
		key string
	)
	// phrase loop
	for _, item := range lang.phraseMap {
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
