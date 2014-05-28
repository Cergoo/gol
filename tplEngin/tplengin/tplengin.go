/*
	tpl & language pkg
	supports multi language output
	(c) 2014 Cergoo
	under terms of ISC license
*/

/*
	! == 33
	" == 34
	> == 62
	{ == 123
	. == 46
	/ == 47
	 == 32
	f == 102
	% == 37
*/

package tplengin

import (
	"fmt"
	"github.com/looplab/tarjan"
	"gol/err"
	"gol/refl"
	"gol/tplEngin/i18n"
	"gol/tplEngin/parser"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// types dictionarys define
type (
	Ttpl struct {
		I18n i18n.Ti18n              // языковой ресурс
		tpl  map[string]*parser.Ttpl // распарсеные шаблоны
		f    refl.FuncMap            // функции (is a extension tags)
	}

	// текс
	tTagText []byte
	// id передаваемого контекста
	tTagVar string
	// extendet tags
	tTagFunc []interface{}
	// id фразы языкового ресурса 0- resurce name, 1- context name type []string
	tTagi18nCon [2]string
	tTagi18nVar [2]string
	// static include temlate
	tTagIncludeCon struct {
		tpl        *parser.Ttpl // распарсеный шаблон
		contextVar string       // id передаваемого контекста
	}
	tTagIncludeVar [2]string
	//
	TReplaser struct {
		i18n *i18n.TReplacer
		tpl  *Ttpl
	}
)

// Create language obj
func New() *Ttpl {
	tpl := new(Ttpl)
	tpl.I18n = i18n.New()
	tpl.tpl = make(map[string]*parser.Ttpl)
	tpl.f = make(refl.FuncMap)
	return tpl
}

func (t *Ttpl) NewReplacer(lang string) *TReplaser {
	r := new(TReplaser)
	r.i18n = t.I18n.NewReplacer(lang)
	r.tpl = t
	return r
}

// Add function to a tpl
func (t *Ttpl) FuncAdd(name string, f interface{}) {
	t.f.Add(name, f)
}

// Loade template
func (t *Ttpl) Load(patch string) {
	var (
		fileSource []byte
	)
	base := "/" + filepath.Base(patch) + "/"

	toparse := new(parser.ToParse)
	toparse.Delimiter[0] = []byte("{{")
	toparse.Delimiter[1] = []byte("}}")
	toparse.ParseTag = parseTag
	toparse.ParseText = parseText

	fileList, e := ioutil.ReadDir(patch)
	err.Panic(e)
	for _, item := range fileList {
		fileSource, e = ioutil.ReadFile(patch + "/" + item.Name())
		t.tpl[base+item.Name()] = parser.Parse(fileSource, toparse)
	}
	fmt.Printf("%v", t.tpl["/maintpls/1.tpl"])
}

//************** control blok begin

// 1. for a tTagInclude: 1)chort name to full name convert for a tTagIncludeCon, 2) prepare context name for a tTagIncludeVar
func (t *Ttpl) tagIncludeCon_ChortToFullName() {
	for key, val := range t.tpl {
		base := "/" + strings.SplitN(key, "/", 3)[1] + "/"
		for key1 := range val.Items {
			switch v := val.Items[key1].(type) {
			case tTagIncludeVar:
				if v[0][0] == 46 {
					// убираем точку
					v[0] = v[0][1:]
				} else if v[0][0] != 47 {
					// короткое имя преобразуем в полное
					v[0] = base + v[0]
					val.Items[key1] = v
				}
			}
		}
	}
}

// 3. for a tTagIncludeCon init before use
func (t *Ttpl) tagIncludeCon_Init() {
	for key, val := range t.tpl {
		for key1 := range val.Items {
			switch v := val.Items[key1].(type) {
			case tTagIncludeVar:
				// не полное имя и не имя переменной контекста - значит короткое имя
				if v[0][0] == 47 {
					tpl := t.tpl[v[0]]
					if tpl == nil {
						err.Panic(err.New("Err parse from tpl: '"+key+"' include tag. Not found tpl: '"+v[0]+"'", 0))
					}
					val.Items[key1] = &tTagIncludeCon{tpl: tpl, contextVar: v[1]}
				}
			}
		}
	}
}

// 2. Контроль зацикливаний тега tTagIncludeCon
func chekloop(tpls map[string]*parser.Ttpl) {
	var (
		key   string
		mitem = []interface{}{}
	)
	matrix := make(map[interface{}][]interface{}, len(tpls))
	for key = range tpls {
		for _, item := range tpls[key].Items {
			switch v := item.(type) {
			case tTagIncludeVar:
				if key == v[0] {
					err.Panic(err.New("Error: loops detection: "+key+" - "+key, 0))
				}
				mitem = append(mitem, v[0])
			}
		}
		if len(mitem) > 0 {
			matrix[key] = mitem
		}
	}
	loop := tarjan.Connections(matrix)
	for i := range loop {
		if len(loop[i]) > 1 {
			err.Panic(err.New("Error: loops detection: "+fmt.Sprint(loop[i]), 0))
		}
	}
}

//
func (t *Ttpl) InitBeforeUse() {
	t.tagIncludeCon_ChortToFullName()
	chekloop(t.tpl)
	t.tagIncludeCon_Init()
}

//************** control blok end

func parseText(source []byte) interface{} {
	return tTagText(source)
}

func parseTag(source []byte) (tag interface{}, contextLen uint16) {
	defer func() {
		if e := recover(); e != nil {
			v := e.(*err.OpenErr)
			v.Text += "err parse tpl: " + string(source)
			err.Panic(v)
		}
	}()

	list := parser.SplitWord(source, 32)
	switch list[0] {
	// основные теги
	case "!":
	case "inc":
		tag = parseTagInclude(list[1:])
	case "i18n":
		tag = parseTagi18n(list[1:])
	default:
		// переменная контекста
		// либо функция (расширенные теги)
		if list[0][0] == 46 {
			tag = tTagVar(list[0][1:])
		} else {
			tag = parseTagFunc(list)
		}
	}
	return
}

// parseTagFunction
func parseTagFunc(source []string) (r tTagFunc) {
	r = make(tTagFunc, 0, len(source))
	for i := range source {
		if source[i][0] == 46 {
			r = append(r, tTagVar(source[i][1:]))
		} else {
			r = append(r, source[i])
		}
	}
	return
}

// parseTagInclude
func parseTagInclude(source []string) (r tTagIncludeVar) {
	copy(r[:], source)
	return
}

// parseTagi18n
func parseTagi18n(source []string) interface{} {
	var m [2]string
	if len(source) > 1 {
		m[1] = source[1]
	}
	if source[0][0] == 46 {
		m[0] = source[0][1:]
		return tTagi18nVar(m)
	}
	m[0] = source[0]
	return tTagi18nCon(m)
}

func (t *TReplaser) Replace(name string, v map[string]interface{}, rezult []byte) []byte {
	tpl := t.tpl.tpl[name]
	if tpl == nil {
		err.Panic(err.New("Err, tpl '"+name+"' not found", 0))
	}
	return t.replace(tpl, v, rezult)
}

func (t *TReplaser) replace(tpl *parser.Ttpl, v map[string]interface{}, rezult []byte) []byte {
	for i := range tpl.Items {
		switch tag := tpl.Items[i].(type) {
		case tTagi18nCon:
			rezult = append(rezult, t.i18n.P(tag[0], v[tag[1]].([]interface{})...)...)
		case tTagi18nVar:
			rezult = append(rezult, t.i18n.P(v[tag[0]].(string), v[tag[1]])...)
		case tTagText:
			rezult = append(rezult, tag...)
		case tTagVar:
			rezult = append(rezult, []byte(fmt.Sprint(v[string(tag)]))...)
		case tTagIncludeCon:
			rezult = append(rezult, t.replace(tag.tpl, v[tag.contextVar].(map[string]interface{}), rezult)...)
		case tTagIncludeVar:
			rezult = append(rezult, t.Replace(tag[0], v[tag[1]].(map[string]interface{}), rezult)...)
		}
	}
	return rezult
}
