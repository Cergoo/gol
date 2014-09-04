// (c) 2014 Cergoo
// under terms of ISC license

// Package humanize it's a pluggable module from i18n pkg
package humanize

import (
	"github.com/Cergoo/gol/err"
	"github.com/Cergoo/gol/reflect/refl"
	"github.com/Cergoo/gol/tplEngin/i18n/human"
	"github.com/Cergoo/gol/tplEngin/i18n/i18n"
)

// Bit mask a functions name to Init
const (
	FHumanByte = 1 << iota
	FHumanBit
	FHumanByteLong
)

// Init module function
func Init(vi18n *i18n.Ti18n, pathToLangResurce string, functions int) {
	vi18n.Load(pathToLangResurce)
	if functions&FHumanBit == FHumanBit {
		vi18n.UserFunc("humanBit", humanBit)
	}
	if functions&FHumanByte == FHumanByte {
		vi18n.UserFunc("humanByte", humanByte)
	}
	if functions&FHumanByteLong == FHumanByteLong {
		vi18n.UserFunc("humanByteLong", humanByteLong)
	}
}

// HumanByte short humanize byte
func humanByte(lang *i18n.Tlang) func([]interface{}) []byte {
	list, ok := lang.Lists["byteshort"]
	err.PanicBool(ok, "i18n list 'byte' not found", 0)
	return human1024(list)
}

// HumanBit short humanize bit
func humanBit(lang *i18n.Tlang) func([]interface{}) []byte {
	list, ok := lang.Lists["bitshort"]
	err.PanicBool(ok, "i18n list 'bit' not found", 0)
	return human1024(list)
}

// HumanByteLong from full name
func humanByteLong(lang *i18n.Tlang) func([]interface{}) []byte {
	list, ok := lang.Lists["+prefix1000"]
	err.PanicBool(ok, "i18n list '+prefix1000' not found", 0)
	pluralByte, ok := lang.Plural["byte"]
	err.PanicBool(ok, "i18n plural 'byte' not found", 0)

	return func(v []interface{}) []byte {
		r, ok := refl.Uint(v[0])
		if !ok {
			return nil
		}
		val, i, valf := human.Byten(r)
		if !(len(list) > int(i)) {
			return nil
		}
		return []byte(val + " " + list[i] + pluralByte[lang.PluralRule(valf)])
	}
}

func human1024(list []string) func([]interface{}) []byte {
	return func(v []interface{}) []byte {
		r, ok := refl.Uint(v[0])
		if !ok {
			return nil
		}
		val, i, _ := human.Byten(r)
		if len(list) > int(i) {
			return []byte(val + list[i])
		}
		return nil
	}
}
