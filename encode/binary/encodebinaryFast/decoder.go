// (c) 2014 Cergoo
// under terms of ISC license

package encodebinaryFast

import (
	//"fmt"
	"github.com/Cergoo/gol/encode/binary/primitive"
	"github.com/Cergoo/gol/err"
	"reflect"
)

// Decode generate decode function from value
func (t *TGen) Decode(val interface{}) {
	valType := reflect.TypeOf(val)
	t.src = "\nfunc Decode(buf IBuf) (t " + typeName(valType.String(), true) + ", e error) {\n"
	t.src += "var (\npart []byte\nbt byte\nln uint32\n)\n\n"
	t.stackName.Push("t")
	t.decode(valType)
	t.src += "return\n}\n\n"
	t.stackName.Clear()
	t.tmpNameGen.Clear()
	_, e := t.f.Write([]byte(t.src))
	err.Panic(e)
}

func (t *TGen) decUint8(name string) {
	t.src += "bt, e = buf.ReadByte()\nif e != nil { return }\n" + name + "= bt\n"
}

func (t *TGen) decUint16(name string) {
	t.src += "part, e = buf.ReadNext(WORD16)\nif e != nil { return }\n" + name + "=Pack.Uint16(part)\n"
}

func (t *TGen) decUint32(name string) {
	t.src += "part, e = buf.ReadNext(WORD32)\nif e != nil { return }\n" + name + "=Pack.Uint32(part)\n"
}

func (t *TGen) decUint64(name string) {
	t.src += "part, e = buf.ReadNext(WORD64)\nif e != nil { return }\n" + name + "=Pack.Uint64(part)\n"
}

func (t *TGen) decUint(name string) {
	t.src += "part, e = buf.ReadNext(WORD64)\nif e != nil { return }\n" + name + "=uint(Pack.Uint64(part))\n"
}

func (t *TGen) decInt8(name string) {
	t.src += "bt, e = buf.ReadByte()\nif e != nil { return }\n" + name + "=int8(bt)\n"
}

func (t *TGen) decInt16(name string) {
	t.src += "part, e = buf.ReadNext(WORD16)\nif e != nil { return }\n" + name + "=int16(Pack.Uint16(part))\n"
}

func (t *TGen) decInt32(name string) {
	t.src += "part, e = buf.ReadNext(WORD32)\nif e != nil { return }\n" + name + "=int32(Pack.Uint32(part))\n"
}

func (t *TGen) decInt64(name string) {
	t.src += "part, e = buf.ReadNext(WORD64)\nif e != nil { return }\n" + name + "=int64(Pack.Uint64(part))\n"
}

func (t *TGen) decInt(name string) {
	t.src += "part, e = buf.ReadNext(WORD64)\nif e != nil { return }\n" + name + "=int(Pack.Uint64(part))\n"
}

func (t *TGen) decFloat32(name string) {
	t.src += "part, e = buf.ReadNext(WORD32)\nif e != nil { return }\n" + name + "=math.Float32frombits(Pack.Uint32(part))\n"
}

func (t *TGen) decFloat64(name string) {
	t.src += "part, e = buf.ReadNext(WORD64)\nif e != nil { return }\n" + name + "=math.Float64frombits(Pack.Uint64(part))\n"
}

func (t *TGen) decComplex64(name string) {
	t.src += "part, e = buf.ReadNext(WORD64)\nif e != nil { return }\n" + name + "=Complex64(part)\n"
}

func (t *TGen) decComplex128(name string) {
	t.src += "part, e = buf.ReadNext(16)\nif e != nil { return }\n" + name + "=Complex128(part)\n"
}

func (t *TGen) decBool(name string) {
	t.src += "bt, e = buf.ReadByte()\n if e != nil { return }\n" + name + "=Bool(bt)\n"
}

func (t *TGen) decString(name string) {
	t.src += "// string decode\n"
	t.src += "part, e = buf.ReadNext(WORD32)\n if e != nil { return }\n ln = Pack.Uint32(part)\n part, e = buf.ReadNext(int(ln))\n if e != nil { return }\n " + name + "=string(part)\n"
}

// Nil
func (t *TGen) decNilBegin(name string) {
	t.src += "bt, e = buf.ReadByte()\n if e != nil { return }\n if bt == 0 { " + name + "=nil } else {\n"
}

func (t *TGen) decNilEnd() {
	t.src += "}\n"
}

func (t *TGen) decPtr(name string, val reflect.Type) {
	t.decNilBegin(name)
	t.src += name + "= new(" + typeName(val.String(), false) + ")\n"
	t.stackName.Push("(*" + name + ")")
	t.decode(val.Elem())
	t.decNilEnd()
}

func (t *TGen) decStruct(name string, val reflect.Type) {
	if val == primitive.TimeType {
		t.src += "part, e = buf.ReadNext(WORD64)\n if e != nil { return }\n" + name + "=Time(part)\n"
	} else {
		var f reflect.StructField
		count := val.NumField()
		for i := 0; i < count; i++ {
			f = val.Field(i)
			if f.PkgPath != "" {
				continue
			}
			t.stackName.Push(name + "." + f.Name)
			t.decode(f.Type)
		}
	}
}

func (t *TGen) decSlice(name string, val reflect.Type) {
	t.src += "// slice decode\n"
	t.decNilBegin(name)
	lnname := t.tmpNameGen.Get()
	idname := t.tmpNameGen.Get()
	t.decUint32(lnname + ":")
	if val.Elem().Kind() == reflect.Uint8 && val.Elem().String() == "uint8" {
		t.src += name + ", e = buf.ReadNext(int(" + lnname + "))\n if e != nil { return }\n"
	} else {
		t.src += name + "=make(" + typeName(val.String(), true) + ", " + lnname + ")\n"
		t.src += "for " + idname + " := uint32(0); " + idname + "<" + lnname + "; " + idname + "++ {\n"
		t.stackName.Push(name + "[" + idname + "]")
		t.decode(val.Elem())
		t.src += "}\n"
	}
	t.decNilEnd()
}

func (t *TGen) decArray(name string, val reflect.Type) {
	t.src += "// array decode\n"
	lnname := t.tmpNameGen.Get()
	idname := t.tmpNameGen.Get()
	t.decUint32(lnname + ":")
	t.src += "for " + idname + " := uint32(0); " + idname + "<" + lnname + "; " + idname + "++ {\n"
	t.stackName.Push(name + "[" + idname + "]")
	t.decode(val.Elem())
	t.src += "}\n"
}

func (t *TGen) decMap(name string, val reflect.Type) {
	t.src += "// map decode\n"
	t.decNilBegin(name)
	idname := t.tmpNameGen.Get()
	lnname := t.tmpNameGen.Get()
	tmpkey := t.tmpNameGen.Get()
	tmpval := t.tmpNameGen.Get()
	t.decUint32(lnname + ":")
	t.src += name + "=make(" + typeName(val.String(), true) + ", " + lnname + ")\n"
	t.src += "var (\n" + tmpkey + " " + typeName(val.Key().String(), true) + "\n" + tmpval + " " + typeName(val.Elem().String(), true) + "\n)\n"
	t.src += "for " + idname + " := uint32(0); " + idname + "<" + lnname + "; " + idname + "++ {\n"
	t.stackName.Push(tmpkey)
	t.decode(val.Key())
	t.stackName.Push(tmpval)
	t.decode(val.Elem())
	t.src += name + "[" + tmpkey + "] = " + tmpval + "\n"
	t.src += "}\n"
	t.decNilEnd()
}

func (t *TGen) decode(val reflect.Type) {
	v, ok := t.stackName.Pop()
	if !ok {
		err.Panic(err.New("stack name empty", 0))
	}
	name := v.(string)

	switch val.Kind() {
	case reflect.Uint8:
		t.decUint8(name)
	case reflect.Uint16:
		t.decUint16(name)
	case reflect.Uint32:
		t.decUint32(name)
	case reflect.Uint64:
		t.decUint64(name)
	case reflect.Uint:
		t.decUint(name)
	case reflect.Int8:
		t.decInt8(name)
	case reflect.Int16:
		t.decInt16(name)
	case reflect.Int32:
		t.decInt32(name)
	case reflect.Int64:
		t.decInt64(name)
	case reflect.Int:
		t.decInt(name)
	case reflect.Float32:
		t.decFloat32(name)
	case reflect.Float64:
		t.decFloat64(name)
	case reflect.Complex64:
		t.decComplex64(name)
	case reflect.Complex128:
		t.decComplex128(name)
	case reflect.Bool:
		t.decBool(name)
	case reflect.String:
		t.decString(name)
	case reflect.Slice:
		t.decSlice(name, val)
	case reflect.Array:
		t.decArray(name, val)
	case reflect.Ptr:
		t.decPtr(name, val)
	case reflect.Struct:
		t.decStruct(name, val)
	case reflect.Map:
		t.decMap(name, val)
	case reflect.Interface:
		err.Panic(err.New("interface type not supported, only strong typed", 0))
	}
}
