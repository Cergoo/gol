// (c) 2014 Cergoo
// under terms of ISC license

package encodejson

func Escape(val []byte) []byte {
	r := make([]byte, len(val))
	for _, v := range val {
		switch v {
		case '"', '\\', '/':
			r = append(r, '\\', v)
		case '\b':
			r = append(r, '\\', 'b')
		case '\f':
			r = append(r, '\\', 'f')
		case '\n':
			r = append(r, '\\', 'n')
		case '\r':
			r = append(r, '\\', 'r')
		case '\t':
			r = append(r, '\\', 't')
		default:
			r = append(r, v)
		}
	}
	return r
}
