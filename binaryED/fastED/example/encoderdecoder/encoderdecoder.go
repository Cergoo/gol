// It's file auto generate

package encoderdecoder

import (
	"github.com/Cergoo/gol/binaryED/fastED/example/exportedtypes"
	. "github.com/Cergoo/gol/binaryED/primitive"
)

func Encode(t *exportedtypes.T1, buf IBuf) {
	if t == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		Pack.PutUint64(buf.Reserve(8), uint64((*t).N1))
		PutString(buf, (*t).N2)
		buf.WriteByte((*t).N3)
		if (*t).N4 == nil {
			buf.WriteByte(0)
		} else {
			buf.WriteByte(1)
			Pack.PutUint64(buf.Reserve(8), uint64((*(*t).N4).N1))
			PutString(buf, (*(*t).N4).N2)
			buf.WriteByte((*(*t).N4).N3)
			// slice encode
			if (*(*t).N4).N4 == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				Pack.PutUint32(buf.Reserve(4), uint32(len((*(*t).N4).N4)))
				for _, tmp1 := range (*(*t).N4).N4 {
					PutString(buf, tmp1)
				}
			}
			// map encode
			if (*(*t).N4).N5 == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				Pack.PutUint32(buf.Reserve(4), uint32(len((*(*t).N4).N5)))
				for tmp2, tmp3 := range (*(*t).N4).N5 {
					Pack.PutUint64(buf.Reserve(8), uint64(tmp2))
					PutString(buf, tmp3)
				}
			}
			// map encode
			if (*(*t).N4).N6 == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				Pack.PutUint32(buf.Reserve(4), uint32(len((*(*t).N4).N6)))
				for tmp4, tmp5 := range (*(*t).N4).N6 {
					Pack.PutUint64(buf.Reserve(8), uint64(tmp4))
					if tmp5 == nil {
						buf.WriteByte(0)
					} else {
						buf.WriteByte(1)
						Pack.PutUint64(buf.Reserve(8), uint64((*tmp5).N))
					}
				}
			}
		}
	}
}

func Decode(buf IBuf) (t *exportedtypes.T1, e error) {
	var (
		part []byte
		bt   byte
		ln   uint32
	)

	bt, e = buf.ReadByte()
	if e != nil {
		return
	}
	if bt == 0 {
		t = nil
	} else {
		if t == nil {
			t = new(exportedtypes.T1)
		}
		part, e = buf.ReadNext(WORD64)
		if e != nil {
			return
		}
		(*t).N1 = int(Pack.Uint64(part))
		// string decode
		part, e = buf.ReadNext(WORD32)
		if e != nil {
			return
		}
		ln = Pack.Uint32(part)
		part, e = buf.ReadNext(int(ln))
		if e != nil {
			return
		}
		(*t).N2 = string(part)
		bt, e = buf.ReadByte()
		if e != nil {
			return
		}
		(*t).N3 = bt
		bt, e = buf.ReadByte()
		if e != nil {
			return
		}
		if bt == 0 {
			(*t).N4 = nil
		} else {
			if (*t).N4 == nil {
				(*t).N4 = new(exportedtypes.T2)
			}
			part, e = buf.ReadNext(WORD64)
			if e != nil {
				return
			}
			(*(*t).N4).N1 = int(Pack.Uint64(part))
			// string decode
			part, e = buf.ReadNext(WORD32)
			if e != nil {
				return
			}
			ln = Pack.Uint32(part)
			part, e = buf.ReadNext(int(ln))
			if e != nil {
				return
			}
			(*(*t).N4).N2 = string(part)
			bt, e = buf.ReadByte()
			if e != nil {
				return
			}
			(*(*t).N4).N3 = bt
			// slice decode
			bt, e = buf.ReadByte()
			if e != nil {
				return
			}
			if bt == 0 {
				(*(*t).N4).N4 = nil
			} else {
				part, e = buf.ReadNext(WORD32)
				if e != nil {
					return
				}
				tmp1 := Pack.Uint32(part)
				(*(*t).N4).N4 = make([]string, tmp1)
				for tmp2 := uint32(0); tmp2 < tmp1; tmp2++ {
					// string decode
					part, e = buf.ReadNext(WORD32)
					if e != nil {
						return
					}
					ln = Pack.Uint32(part)
					part, e = buf.ReadNext(int(ln))
					if e != nil {
						return
					}
					(*(*t).N4).N4[tmp2] = string(part)
				}
			}
			// map decode
			bt, e = buf.ReadByte()
			if e != nil {
				return
			}
			if bt == 0 {
				(*(*t).N4).N5 = nil
			} else {
				part, e = buf.ReadNext(WORD32)
				if e != nil {
					return
				}
				tmp4 := Pack.Uint32(part)
				(*(*t).N4).N5 = make(map[int]string, tmp4)
				var (
					tmp5 int
					tmp6 string
				)
				for tmp3 := uint32(0); tmp3 < tmp4; tmp3++ {
					part, e = buf.ReadNext(WORD64)
					if e != nil {
						return
					}
					tmp5 = int(Pack.Uint64(part))
					// string decode
					part, e = buf.ReadNext(WORD32)
					if e != nil {
						return
					}
					ln = Pack.Uint32(part)
					part, e = buf.ReadNext(int(ln))
					if e != nil {
						return
					}
					tmp6 = string(part)
					(*(*t).N4).N5[tmp5] = tmp6
				}
			}
			// map decode
			bt, e = buf.ReadByte()
			if e != nil {
				return
			}
			if bt == 0 {
				(*(*t).N4).N6 = nil
			} else {
				part, e = buf.ReadNext(WORD32)
				if e != nil {
					return
				}
				tmp8 := Pack.Uint32(part)
				(*(*t).N4).N6 = make(map[int]*exportedtypes.T3, tmp8)
				var (
					tmp9  int
					tmp10 *exportedtypes.T3
				)
				for tmp7 := uint32(0); tmp7 < tmp8; tmp7++ {
					part, e = buf.ReadNext(WORD64)
					if e != nil {
						return
					}
					tmp9 = int(Pack.Uint64(part))
					bt, e = buf.ReadByte()
					if e != nil {
						return
					}
					if bt == 0 {
						tmp10 = nil
					} else {
						if tmp10 == nil {
							tmp10 = new(exportedtypes.T3)
						}
						part, e = buf.ReadNext(WORD64)
						if e != nil {
							return
						}
						(*tmp10).N = int(Pack.Uint64(part))
					}
					(*(*t).N4).N6[tmp9] = tmp10
				}
			}
		}
	}
	return
}
