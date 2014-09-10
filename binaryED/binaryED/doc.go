// (c) 2014 Cergoo
// under terms of ISC license

/*
Package binaryED it's a binary structure less format encode/decode implementation
fork github.com/youtube/vitess/go/bson

Attention!

======================================================================

Before you can use this package need to patch standard library reflect,
for unto this add the file the following function:
go/src/pkg/reflect/value.go
    // the Go a user hack
    func (v Value) Ptr() unsafe.Pointer {
       return v.ptr
    }

The package is designed for fast serialization / deserialization:
	uint8, uint16, uint32, uint64, uint
	int8, int16, int32, int64, int
	floate32, floate64
	bool
	string
	slise, array
	map (keys not pointer type)
	struct

Description of the coding format

Binary encoding of the data on the basis of binary.LittleEndian.
Before the reference types: slices, maps, pointers, added byte: 0-nil, 1-not nil.
Before the strings, arrays, slices, maps, added uint32 number of items.

Important !

======================================================================

- Nonexported field structures are ignored.

- In decoding the variable structure is used in which the decoding occurs,
necessary to match the structure of the receiver structure of the source
up to the order of the fields in the description of the structures.

- Possible encoding / decoding only a strictly structured data,
ie map[string]interfase {} can not be coded as values ​​map do not have a strict structure.

Важно!

- При кодировании/декодировании неэкспортируемые поля структур игнорируются.

- При декодировании используется структура переменной в которую происходит декодирование, необходимо чтобы структура приёмнника соответсвовала структуре источника вплоть до порядка следования полей в описании структур.

- Возможно кодирование/декодирование только строго структурированных данных,
то есть map[string]interfase{} нельзя кодировать так как значения хештаблицы
не имеют описания структуры а формат кодирования некодирует описание структур а только их данные.

Описание формата кодирования

Бинарное кодирование данных на основе binary.LittleEndian.
Перед ссылочными типами: срезами, отображениями, указателями, добавляется байт: 0-nil, 1-not nil.
Перед строками, массивами, срезами, отображениями, добавляется uint32 количество элементов.
*/
package binaryED
