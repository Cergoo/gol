// (c) 2014 Cergoo
// under terms of ISC license

package fastED

import (
	//"fmt"
	"github.com/Cergoo/gol/binaryED/primitive"
	"github.com/Cergoo/gol/binaryED/tmpName"
	"github.com/Cergoo/gol/err"
	extfilepath "github.com/Cergoo/gol/filepath"
	"github.com/Cergoo/gol/stack/stack"
	"os"
	"path/filepath"
	"reflect"
)

type (
	TGen struct {
		stackName  stack.TStack
		tmpNameGen *tmpName.TtmpName
		src        string
		f          *os.File
	}
)

func New(outputFile string, imported ...string) *TGen {
	os.Remove(outputFile)
	f, e := os.Create(outputFile)
	err.Panic(e)
	outputFile = filepath.Base(outputFile)
	outputFile, _ = extfilepath.Ext(outputFile)
	f.Write([]byte("// It's file auto generate \n\n"))
	f.Write([]byte("package " + outputFile + "\n\n"))

	resultImport := ". \"github.com/Cergoo/gol/binaryED/primitive\"\n"
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

func (t *TGen) Close() {
	t.f.Close()
}

// Encode encode value to binary
func (t *TGen) Encode(val interface{}) {
	valType := reflect.TypeOf(val)
	t.src = "\nfunc Encode(t " + typeName(valType.String(), true) + ", buf IBuf) {\n"
	t.stackName.Push("t")
	t.encode(valType)
	t.src += "}\n\n"
	t.stackName.Clear()
	t.tmpNameGen.Clear()
	_, e := t.f.Write([]byte(t.src))
	err.Panic(e)
}

func (t *TGen) genUint8(name string) {
	t.src += "buf.WriteByte(" + name + ")\n"
}

func (t *TGen) genUint16(name string) {
	t.src += "Pack.PutUint16(buf.Reserve(2), " + name + ")\n"
}

func (t *TGen) genUint32(name string) {
	t.src += "Pack.PutUint32(buf.Reserve(4), " + name + ")\n"
}

func (t *TGen) genUint64(name string) {
	t.src += "Pack.PutUint64(buf.Reserve(8), " + name + ")\n"
}

func (t *TGen) genInt8(name string) {
	t.src += "buf.WriteByte(uint8(" + name + "))\n"
}

func (t *TGen) genInt16(name string) {
	t.src += "Pack.PutUint16(buf.Reserve(2), uint16(" + name + "))\n"
}

func (t *TGen) genInt32(name string) {
	t.src += "Pack.PutUint32(buf.Reserve(4), uint32(" + name + "))\n"
}

func (t *TGen) genInt64(name string) {
	t.src += "Pack.PutUint64(buf.Reserve(8), uint64(" + name + "))\n"
}

func (t *TGen) genFloat32(name string) {
	t.src += "Pack.PutUint32(buf.Reserve(4), math.Float32bits(" + name + "))\n"
}

func (t *TGen) genFloat64(name string) {
	t.src += "Pack.PutUint64(buf.Reserve(8), math.Float64bits(" + name + "))\n"
}

func (t *TGen) genComplex64(name string) {
	t.src += "PutComplex64(buf, " + name + ")\n"
}

func (t *TGen) genComplex128(name string) {
	t.src += "PutComplex128(buf, " + name + ")\n"
}

func (t *TGen) genBool(name string) {
	t.src += "PutBool(buf, " + name + ")\n"
}

func (t *TGen) genString(name string) {
	t.src += "PutString(buf, " + name + ")\n"
}

// Nil
func (t *TGen) genNilBegin(name string) {
	t.src += "if " + name + " == nil { buf.WriteByte(0) } else { \nbuf.WriteByte(1)\n"
}

func (t *TGen) genNilEnd() {
	t.src += "}\n"
}

func (t *TGen) genPtr(name string, val reflect.Type) {
	t.genNilBegin(name)
	t.stackName.Push("(*" + name + ")")
	t.encode(val.Elem())
	t.genNilEnd()
}

func (t *TGen) genStruct(name string, val reflect.Type) {
	if val == primitive.TimeType {
		t.src += "PutTime(buf, " + name + ")\n"
	} else {
		var f reflect.StructField
		count := val.NumField()
		for i := 0; i < count; i++ {
			f = val.Field(i)
			if f.PkgPath != "" {
				continue
			}
			t.stackName.Push(name + "." + f.Name)
			t.encode(f.Type)
		}
	}
}

func (t *TGen) genSlice(name string, val reflect.Type) {
	t.src += "// slice encode\n"
	t.genNilBegin(name)
	t.genUint32("uint32(len(" + name + "))")
	if val.Elem().Kind() == reflect.Uint8 {
		if val.Elem().String() != "uint8" {
			t.src += "buf.Write(*(*[]byte)(unsafe.Pointer(&" + name + ")))\n"
		} else {
			t.src += "buf.Write(" + name + ")\n"
		}
		return
	}
	tmp := t.tmpNameGen.Get() // get tmp var name
	t.src += "for _, " + tmp + " := range " + name + " {\n"
	t.stackName.Push(tmp)
	t.encode(val.Elem())
	t.src += "}\n"
	t.genNilEnd()
}

func (t *TGen) genArray(name string, val reflect.Type) {
	t.src += "// array encode\n"
	t.genUint32("uint32(len(" + name + "))")
	if val.Elem().Kind() == reflect.Uint8 {
		if val.Elem().String() != "uint8" {
			tmp := t.tmpNameGen.Get() // get tmp var name
			t.src += tmp + " := " + name + "[:]\n"
			t.src += "buf.Write(*(*[]byte)(unsafe.Pointer(&" + tmp + ")))\n"
		} else {
			t.src += "buf.Write(" + name + "[:])\n"
		}
		return
	}
	tmp := t.tmpNameGen.Get() // get tmp var name
	t.src += "for _, " + tmp + " := range " + name + " {\n"
	t.stackName.Push(tmp)
	t.encode(val.Elem())
	t.src += "}\n"
}

func (t *TGen) genMap(name string, val reflect.Type) {
	t.src += "// map encode\n"
	t.genNilBegin(name)
	t.genUint32("uint32(len(" + name + "))")
	tmpKey := t.tmpNameGen.Get() // get tmp var name
	tmpVal := t.tmpNameGen.Get() // get tmp var name
	t.src += "for " + tmpKey + ", " + tmpVal + " := range " + name + " {\n"
	t.stackName.Push(tmpKey)
	t.encode(val.Key())
	t.stackName.Push(tmpVal)
	t.encode(val.Elem())
	t.src += "}\n"
	t.genNilEnd()
}

func (t *TGen) encode(val reflect.Type) {
	v, ok := t.stackName.Pop()
	if !ok {
		err.Panic(err.New("stack name empty", 0))
	}
	name := v.(string)

	switch val.Kind() {
	case reflect.Uint8:
		t.genUint8(name)
	case reflect.Uint16:
		t.genUint16(name)
	case reflect.Uint32:
		t.genUint32(name)
	case reflect.Uint64, reflect.Uint:
		t.genUint64(name)
	case reflect.Int8:
		t.genInt8(name)
	case reflect.Int16:
		t.genInt16(name)
	case reflect.Int32:
		t.genInt32(name)
	case reflect.Int64, reflect.Int:
		t.genInt64(name)
	case reflect.Float32:
		t.genFloat32(name)
	case reflect.Float64:
		t.genFloat64(name)
	case reflect.Complex64:
		t.genComplex64(name)
	case reflect.Complex128:
		t.genComplex128(name)
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
		err.Panic(err.New("interface type not supported, only strong typed", 0))
	}
}
