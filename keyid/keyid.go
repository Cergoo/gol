/*
	key to id and id to key association pack
	(c) 2014 Cergoo
	under terms of ISC license
*/
package keyid

type (
	Tkeyid struct {
		keytoid map[string]uint
		idtokey map[uint]string
	}
)

// Constructor
func New() *Tkeyid {
	return &Tkeyid{keytoid: make(map[string]uint), idtokey: make(map[uint]string)}
}

// Set cortege (key, id)
func (t *Tkeyid) Set(key string, id uint) {
	t.idtokey[id] = key
	t.keytoid[key] = id
}

// Delete from key and return id
func (t *Tkeyid) DelFromKey(key string) (id uint, ok bool) {
	id, ok = t.keytoid[key]
	if ok {
		delete(t.idtokey, id)
		delete(t.keytoid, key)
	}
	return
}

// Delete from id and return key
func (t *Tkeyid) DelFromId(id uint) (key string, ok bool) {
	key, ok = t.idtokey[id]
	if ok {
		delete(t.idtokey, id)
		delete(t.keytoid, key)
	}
	return
}

// Get id from key
func (t *Tkeyid) GetId(key string) (id uint, ok bool) {
	id, ok = t.keytoid[key]
	return
}

// Get key from id
func (t *Tkeyid) GetKey(id uint) (key string, ok bool) {
	key, ok = t.idtokey[id]
	return
}
