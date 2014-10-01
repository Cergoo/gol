// (c) 2014 Cergoo
// under terms of ISC license

package encodejsonFast

import (
	//"fmt"
	. "github.com/Cergoo/gol/encode/fastutil"
	. "github.com/Cergoo/gol/encode/json/common"
	"github.com/Cergoo/gol/err"
	extfilepath "github.com/Cergoo/gol/filepath"
	"github.com/Cergoo/gol/stack/stack"
	"github.com/Cergoo/gol/tmpName"
	"os"
	"path/filepath"
	"reflect"
)

type (
	// TGen it's a main struct
	TGen struct {
		stackName  stack.TStack
		tmpNameGen *tmpName.TtmpName
		src        string
		f          *os.File
	}
)

// New create new generator
func New(outputFile string, imported ...string) *TGen {
	os.Remove(outputFile)
	f, e := os.Create(outputFile)
	err.Panic(e)
	outputFile = filepath.Base(outputFile)
	outputFile, _ = extfilepath.Ext(outputFile)
	f.Write([]byte("// It's file auto generate encodejsonFast\n\n"))
	f.Write([]byte("package " + outputFile + "\n\n"))

	resultImport := ". \"github.com/Cergoo/gol/encode/json/common\"\n"
	resultImport += "\"strconv\"\n"
	for i := range imported {
		resultImport += "\"" + imported[i] + "\"\n"
	}
	f.Write([]byte("import (\n" + resultImport + ")\n"))

	return &TGen{
		stackName:  make(stack.TStack, 0, 10),
		tmpNameGen: tmpName.New(),
		f:          f,
	}
}

// Close end work generator
func (t *TGen) Close() {
	t.f.Close()
}

// Encode generate encode function from value
func (t *TGen) Encode(val interface{}) {
	valType := reflect.TypeOf(val)
	t.src = "\nfunc Encode(buf []byte, t " + TypeName(valType.String(), true) + ") []byte {\n"
	t.stackName.Push("t")
	t.encode(valType)
	t.src += "return buf \n}\n\n"
	t.stackName.Clear()
	t.tmpNameGen.Clear()
	_, e := t.f.Write([]byte(t.src))
	err.Panic(e)
}

// Nil
func (t *TGen) genNilBegin(name string) {
	t.src += "if " + name + " == nil { buf=append(buf, Null...) } else {\n"
}

func (t *TGen) genNilEnd() {
	t.src += "}\n"
}

func (t *TGen) genUint(name string) {
	t.src += "buf = strconv.AppendUint(buf, uint64(" + name + "), 10)\n"
}

func (t *TGen) genInt(name string) {
	t.src += "buf = strconv.AppendInt(buf, int64(" + name + "), 10)\n"
}

func (t *TGen) genFloat(name string) {
	t.src += "buf = strconv.AppendFloat(buf, float64(" + name + "), 'd', -1, 64)\n"
}

func (t *TGen) genBool(name string) {
	t.src += "if val.Bool() {\nbuf = append(buf, Tru...)\n} else {\nbuf = append(buf, Fal...)\n}"
}

func (t *TGen) genString(name string) {
	t.src += "buf = WriteJsonString(buf, []byte(" + name + "))\n"
}

func (t *TGen) genPtr(name string, val reflect.Type) {
	t.genNilBegin(name)
	t.stackName.Push("(*" + name + ")")
	t.encode(val.Elem())
	t.genNilEnd()
}

func (t *TGen) genInterface(name string, val reflect.Type) {
	t.genNilBegin(name)
	t.stackName.Push(name)
	t.encode(val.Elem())
	t.genNilEnd()
}

func (t *TGen) genStruct(name string, val reflect.Type) {
	t.src += "buf = append(buf, '{')\n"
	var f reflect.StructField
	ln := val.NumField()
	for i := 0; i < ln; i++ {
		f = val.Field(i)
		if f.PkgPath != "" {
			continue
		}

		var tmp []byte
		for _, n := range f.Name[:] {
			tmp = append(tmp, '\'', byte(n), '\'', ',')
		}
		t.src += "buf = append(buf, '\"'," + string(tmp) + "'\"',':')\n"
		t.stackName.Push(name + "." + f.Name)
		t.encode(f.Type)
		if i < ln-1 {
			t.src += "buf = append(buf, ',')\n"
		}
	}
	t.src += "buf = append(buf, '}')\n"
}

func (t *TGen) genSlice(name string, val reflect.Type) {
	t.src += "// slice encode\n"
	if val.Elem().Kind() == reflect.Uint8 {
		if val.Elem().String() != "uint8" {
			t.src += "WriteJsonString(buf, *(*[]byte)(unsafe.Pointer(&" + name + ")))\n"
		} else {
			t.src += "WriteJsonString(buf, " + name + ")\n"
		}
		return
	}

	t.genNilBegin(name)
	t.src += "if len(" + name + ")>0 {\n"
	t.src += "buf = append(buf, '[')\n"
	tmp := t.tmpNameGen.Get() // get tmp var name
	t.src += "for _, " + tmp + " := range " + name + " {\n"
	t.stackName.Push(tmp)
	t.encode(val.Elem())
	t.src += "buf = append(buf, ',')\n"
	t.src += "}\n"
	t.src += "buf[len(buf)-1] = ']'\n"
	t.src += "} else {\nbuf = append(buf, '[', ']')\n}\n"
	t.genNilEnd()
}

func (t *TGen) genArray(name string, val reflect.Type) {
	t.src += "// array encode\n"
	if val.Elem().Kind() == reflect.Uint8 {
		if val.Elem().String() != "uint8" {
			t.src += "WriteJsonString(buf, *(*[]byte)(unsafe.Pointer(&" + name + "[:])))\n"
		} else {
			t.src += "WriteJsonString(buf, " + name + "[:])\n"
		}
		return
	}

	t.src += "if len(" + name + ")>0 {\n"
	t.src += "buf = append(buf, '[')\n"
	tmp := t.tmpNameGen.Get() // get tmp var name
	t.src += "for _, " + tmp + " := range " + name + " {\n"
	t.stackName.Push(tmp)
	t.encode(val.Elem())
	t.src += "buf = append(buf, ',')\n"
	t.src += "}\n"
	t.src += "buf[len(buf)-1] = ']'\n"
	t.src += "} else {\nbuf = append(buf, '[', ']')\n}\n"
}

func (t *TGen) genMap(name string, val reflect.Type) {
	t.src += "// map encode\n"
	t.genNilBegin(name)

	tmpKey := t.tmpNameGen.Get() // get tmp var name
	tmpVal := t.tmpNameGen.Get() // get tmp var name
	// as object
	if val.Key().Kind() == reflect.String {
		t.src += "if len(" + name + ")>0 {\n"
		t.src += "buf = append(buf, '{')\n"
		t.src += "for " + tmpKey + ", " + tmpVal + " := range " + name + " {\n"
		t.src += "WriteJsonString(buf, " + tmpKey + "[:])\n"
		t.src += "buf = append(buf, ':')\n"
		t.stackName.Push(tmpVal)
		t.encode(val.Elem())
		t.src += "buf = append(buf, ',')\n"
		t.src += "}\n"
		t.src += "buf[len(buf)-1] = '}'\n"
		t.src += "} else {\nbuf = append(buf, '{', '}')\n}\n"
	} else {
		// as array
		t.src += "if len(" + name + ")>0 {\n"
		t.src += "buf = append(buf, '[')\n"
		t.src += "for " + tmpKey + ", " + tmpVal + " := range " + name + " {\n"
		t.stackName.Push(tmpKey)
		t.encode(val.Key())
		t.src += "buf = append(buf, ',')\n"
		t.stackName.Push(tmpVal)
		t.encode(val.Elem())
		t.src += "buf = append(buf, ',')\n"
		t.src += "}\n"
		t.src += "buf[len(buf)-1] = ']'\n"
		t.src += "} else {\nbuf = append(buf, '[', ']')\n}\n"
	}

	t.genNilEnd()
}

func (t *TGen) encode(val reflect.Type) {
	v, ok := t.stackName.Pop()
	if !ok {
		err.Panic(err.New("stack name empty", 0))
	}
	name := v.(string)

	// use json.Marshaler implement
	if val.Implements(MarshalerType) {
		tmpName := t.tmpNameGen.Get() // get tmp var name
		t.src += "tmpName, _ := " + name + ".MarshalJSON()\n"
		t.src += "buf = append(buf, " + tmpName + ")\n"
		return
	}

	switch val.Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		t.genUint(name)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		t.genInt(name)
	case reflect.Float32, reflect.Float64:
		t.genFloat(name)
	case reflect.Bool:
		t.genBool(name)
	case reflect.String:
		t.genString(name)
	case reflect.Slice:
		t.genSlice(name, val)
	case reflect.Array:
		t.genArray(name, val)
	case reflect.Ptr:
		t.genPtr(name, val)
	case reflect.Struct:
		t.genStruct(name, val)
	case reflect.Map:
		t.genMap(name, val)
	case reflect.Interface:
		t.genInterface(name, val)
	}
}
