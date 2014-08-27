// (c) 2014 Cergoo
// under terms of ISC license

// Package keyid it's key to id and id to key association implementation
package keyid

type (
  // Tkeyid struct 
	Tkeyid struct {
		keytoid map[string]uint
		idtokey map[uint]string
	}
)

// New constructor of a new Tkeyid 
func New() *Tkeyid {
	return &Tkeyid{keytoid: make(map[string]uint), idtokey: make(map[uint]string)}
}

// Set set cortege (key, id)
func (t *Tkeyid) Set(key string, id uint) {
	t.idtokey[id] = key
	t.keytoid[key] = id
}

// DelFromKey delete from key and return id
func (t *Tkeyid) DelFromKey(key string) (id uint, ok bool) {
	id, ok = t.keytoid[key]
	if ok {
		delete(t.idtokey, id)
		delete(t.keytoid, key)
	}
	return
}

// DelFromId delete from id and return key
func (t *Tkeyid) DelFromId(id uint) (key string, ok bool) {
	key, ok = t.idtokey[id]
	if ok {
		delete(t.idtokey, id)
		delete(t.keytoid, key)
	}
	return
}

// GetId get id from key
func (t *Tkeyid) GetId(key string) (id uint, ok bool) {
	id, ok = t.keytoid[key]
	return
}

// GetKey get key from id
func (t *Tkeyid) GetKey(id uint) (key string, ok bool) {
	key, ok = t.idtokey[id]
	return
}
