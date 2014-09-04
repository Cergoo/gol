// (c) 2014 Cergoo
// under terms of ISC license

package i18n

import (
	"fmt"
	"github.com/Cergoo/gol/err"
	gfilepath "github.com/Cergoo/gol/filepath"
	"github.com/Cergoo/gol/jsonConfig"
	"github.com/Cergoo/gol/reflect/refl"
	"github.com/Cergoo/gol/tplEngin/i18n/plural"
	"github.com/Cergoo/gol/tplEngin/parser"
	"github.com/Cergoo/gol/util"
	// "github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

// types dictionarys define
type (
	Tlang struct {
		PluralRule plural.PluralRule                     // language plural rule
		Plural     map[string][]string                   // plural pronunciation
		items      map[string]*titem                     // tpl
		Lists      map[string][]string                   // other lists
		F          map[string]func([]interface{}) []byte // user plugin functions
	}
	titem struct {
		items        parser.Ttpl // tags
		contextCount int         // context var count
	}

	tTagText []byte
	tTagVar  struct {
		id     uint16
		format string
	}
	tTagPlural struct {
		count uint16
		text  []string
	}
	tTagFunction struct {
		fname string
		f     func([]interface{}) []byte
		vars  []uint16
	}

	// Ti18n it's a main struct
	Ti18n struct {
		lang map[string]*Tlang
	}
	// TReplacer replacer template to current lang words from i18n
	TReplacer struct {
		langName string
		lang     *Tlang
	}
)

// UserFunc registry user pluggable function into i18n
func (t *Ti18n) UserFunc(name string, f func(lang *Tlang) func([]interface{}) []byte) {
	for _, v := range t.lang {
		v.F[name] = f(v)
	}
}

// NewReplacer Create new replacer from language resources
func (t *Ti18n) NewReplacer(langName string) (*TReplacer, error) {
	lang, e := t.lang[langName]
	if !e {
		return nil, err.New("Not found lang resurse from langname: '"+langName+"'", 0)
	}
	return &TReplacer{langName: langName, lang: lang}, nil
}

// Lang get lang
func (t *TReplacer) Lang() string {
	return t.langName
}

// P print phrase
func (t *TReplacer) P(key string, context ...interface{}) []byte {
	tpl := t.lang.items[key]
	if tpl == nil {
		return nil
	}
	return t.p(tpl, context)
}

// Plural get plural word form. Use if Load (pluralAccess)
func (t *TReplacer) Plural(key string, count float64) string {
	v, e := t.lang.Plural[key]
	if e {
		return v[t.lang.PluralRule(count)]
	}
	return ""
}

// New create new language object
func New() *Ti18n {
	return &Ti18n{
		lang: make(map[string]*Tlang),
	}
}

// Init run it's after all loads lang resurce
func (t *Ti18n) Init(pluralAccess bool) {
	for key, val := range t.lang {
		initAfterParse(val, key)
		if !pluralAccess {
			val.Plural = nil
		}
	}
}

// Load load language resources
func (t *Ti18n) Load(patch string) {
	type (
		tmpLang struct {
			PluralRule string
			Plural     map[string][]string
			Phrase     map[string]string
			Lists      map[string][]string
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
		vtmpLang.Lists = make(map[string][]string)
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
		lang := &Tlang{
			items:      make(map[string]*titem),
			Plural:     item.Plural,
			Lists:      item.Lists,
			PluralRule: plural.PluralRules[item.PluralRule],
			F:          make(map[string]func([]interface{}) []byte),
		}
		if lang.PluralRule == nil && len(lang.Plural) > 0 {
			err.Panic(err.New("Not found plural rule: '"+item.PluralRule+"'", 0))
		}

		for keyPhrase, itemPhrase := range item.Phrase {
			lang.items[keyPhrase] = &titem{items: parser.Parse([]byte(itemPhrase), toparse), contextCount: -1}
		}

		existLang := t.lang[key]
		if existLang == nil {
			t.lang[key] = lang
		} else {
			// add phrase
			for key, val := range lang.items {
				existLang.items[key] = val
			}
			// add plural
			for key, val := range lang.Plural {
				existLang.Plural[key] = val
			}
			// add lists
			for key, val := range lang.Lists {
				existLang.Lists[key] = val
			}
		}
	}
}

// Get phrase
func (t *TReplacer) p(tpl *titem, context []interface{}) []byte {
	var (
		varFloate float64
		varBool   bool
		varUint16 uint16
		result    []byte
	)
	if int(tpl.contextCount) > len(context) {
		err.Panic(err.New("i18n Mismatch context len: ("+strconv.Itoa(int(tpl.contextCount))+" , "+strconv.Itoa(len(context))+") "+fmt.Sprintf("%#v", tpl.items), 0))
	}

	for _, item := range tpl.items {
		switch v := item.(type) {
		case tTagText:
			result = append(result, v...)
		case *tTagVar:
			if len(v.format) > 0 {
				result = append(result, []byte(fmt.Sprintf(v.format, context[v.id]))...)
			} else {
				result = append(result, []byte(fmt.Sprint(context[v.id]))...)
			}
		case *tTagPlural:
			varFloate, varBool = refl.Floate(context[v.count])
			if varBool {
				result = append(result, []byte(v.text[t.lang.PluralRule(varFloate)])...)
			}
		case *tTagFunction:
			var vars []interface{}
			for _, varUint16 = range v.vars {
				vars = append(vars, context[varUint16])
			}
			result = append(result, v.f(vars)...)
		}
	}
	return result
}

// Init tags
func initAfterParse(lang *Tlang, name string) {
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
				util.MaxSet(&item.contextCount, int(v.count))
				v.text, e = lang.Plural[key]
				err.PanicBool(e, "Err parse:"+name+" Not found plural key: "+key, 0)
			case tTagVar:
				util.MaxSet(&item.contextCount, int(v.id))
			case *tTagFunction:
				v.f, e = lang.F[v.fname]
				err.PanicBool(e, "Err parse:"+name+" Not found users function: "+v.fname, 0)
				for _, val := range v.vars {
					util.MaxSet(&item.contextCount, int(val))
				}
			}
		}
		// initially -1 because of len(n) = max(id)+1
		item.contextCount++
	}
}

// parse plural tag
func parseTagPlural(source []string) *tTagPlural {
	if len(source) < 2 {
		err.Panic(err.New("error parsing to Plural Tag: "+fmt.Sprint(source), 0))
	}
	i, e := strconv.ParseUint(source[1], 10, 16)
	err.Panic(e)
	return &tTagPlural{uint16(i), []string{source[0]}}
}

// parse context var tag
func parseTagVar(source []string) (v *tTagVar) {
	i, e := strconv.ParseUint(source[0], 10, 16)
	err.Panic(e)
	v = &tTagVar{id: uint16(i)}
	if len(source) > 1 {
		v.format = source[1]
	}
	return
}

// parse user functions tag
func parseTagFunc(source []string) (v *tTagFunction) {
	var (
		i uint64
		e error
	)
	v = &tTagFunction{fname: source[0]}
	for _, val := range source[1:] {
		i, e = strconv.ParseUint(val, 10, 16)
		err.Panic(e)
		v.vars = append(v.vars, uint16(i))
	}
	return
}

func parseText(source []byte) interface{} {
	return tTagText(source)
}

func parseTag(source []byte) interface{} {
	defer func() {
		if e := recover(); e != nil {
			panic(err.New("err parse i18n: "+string(source), 0))
		}
	}()

	list := parser.SplitWord(source, 32)
	switch list[0] {
	case "plural":
		return parseTagPlural(list[1:])
	case "f":
		return parseTagFunc(list[1:])
	default:
		return parseTagVar(list)
	}
}
