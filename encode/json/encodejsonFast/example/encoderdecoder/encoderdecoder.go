// It's file auto generate encodejsonFast

package encoderdecoder

import (
	. "github.com/Cergoo/gol/encode/json/common"
	"github.com/Cergoo/gol/encode/json/encodejsonFast/example/exportedtypes"
	"strconv"
)

func Encode(buf []byte, t *exportedtypes.T1) []byte {
	if t == nil {
		buf = append(buf, Null...)
	} else {
		buf = append(buf, '{')
		buf = append(buf, '"', 'N', '1', '"', ':')
		buf = strconv.AppendInt(buf, int64((*t).N1), 10)
		buf = append(buf, ',')
		buf = append(buf, '"', 'N', '2', '"', ':')
		buf = WriteJsonString(buf, []byte((*t).N2))
		buf = append(buf, ',')
		buf = append(buf, '"', 'N', '3', '"', ':')
		buf = strconv.AppendUint(buf, uint64((*t).N3), 10)
		buf = append(buf, ',')
		buf = append(buf, '"', 'N', '4', '"', ':')
		if (*t).N4 == nil {
			buf = append(buf, Null...)
		} else {
			buf = append(buf, '{')
			buf = append(buf, '"', 'N', '1', '"', ':')
			buf = strconv.AppendInt(buf, int64((*(*t).N4).N1), 10)
			buf = append(buf, ',')
			buf = append(buf, '"', 'N', '2', '"', ':')
			buf = WriteJsonString(buf, []byte((*(*t).N4).N2))
			buf = append(buf, ',')
			buf = append(buf, '"', 'N', '3', '"', ':')
			buf = strconv.AppendUint(buf, uint64((*(*t).N4).N3), 10)
			buf = append(buf, ',')
			buf = append(buf, '"', 'N', '4', '"', ':')
			// slice encode
			if (*(*t).N4).N4 == nil {
				buf = append(buf, Null...)
			} else {
				if len((*(*t).N4).N4) > 0 {
					buf = append(buf, '[')
					for _, tmp1 := range (*(*t).N4).N4 {
						buf = WriteJsonString(buf, []byte(tmp1))
						buf = append(buf, ',')
					}
					buf[len(buf)-1] = ']'
				} else {
					buf = append(buf, '[', ']')
				}
			}
			buf = append(buf, ',')
			buf = append(buf, '"', 'N', '5', '"', ':')
			// map encode
			if (*(*t).N4).N5 == nil {
				buf = append(buf, Null...)
			} else {
				if len((*(*t).N4).N5) > 0 {
					buf = append(buf, '[')
					for tmp2, tmp3 := range (*(*t).N4).N5 {
						buf = strconv.AppendInt(buf, int64(tmp2), 10)
						buf = append(buf, ',')
						buf = WriteJsonString(buf, []byte(tmp3))
						buf = append(buf, ',')
					}
					buf[len(buf)-1] = ']'
				} else {
					buf = append(buf, '[', ']')
				}
			}
			buf = append(buf, ',')
			buf = append(buf, '"', 'N', '6', '"', ':')
			// map encode
			if (*(*t).N4).N6 == nil {
				buf = append(buf, Null...)
			} else {
				if len((*(*t).N4).N6) > 0 {
					buf = append(buf, '[')
					for tmp4, tmp5 := range (*(*t).N4).N6 {
						buf = strconv.AppendInt(buf, int64(tmp4), 10)
						buf = append(buf, ',')
						if tmp5 == nil {
							buf = append(buf, Null...)
						} else {
							buf = append(buf, '{')
							buf = append(buf, '"', 'N', '"', ':')
							buf = strconv.AppendInt(buf, int64((*tmp5).N), 10)
							buf = append(buf, '}')
						}
						buf = append(buf, ',')
					}
					buf[len(buf)-1] = ']'
				} else {
					buf = append(buf, '[', ']')
				}
			}
			buf = append(buf, '}')
		}
		buf = append(buf, '}')
	}
	return buf
}
