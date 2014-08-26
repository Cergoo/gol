# gol (Go Library)
(c) 2013-2014 Cergoo   
under terms of ISC license

## binaryED
binary Encode Decode implementation pkg

## cache
Cache is an in-memory key:value store/cache similar to memcached that is suitable for applications running on a single machine.    
The package is the implementation of hashmap for organizing in-memory key-value data store. Automatic lifetime management of records can be enabled or disabled. LRU caches are affected by the problem leaching records in intensive add, i.e. the records  permanently pushed do not linger in the cach. This package does not implement the LRU. In this implementation the time life records indicated for all the generated cache, specified time value is the size of time interval during which a new record is guaranteed to live in the cache. Then have a record of lives at least one time interval maximum of two time interval + can be implemented for the "if it's read then it lives" if a record is requested that her life is prolonged for the next time interval.  
    
### Feature:
- thread-safe
- faste and high availability
- increment/decrement command supported
- save/load operation supported
- mechanism of managing the lifetime data: time expirations (set for the entire cache) and options "if read then life", support callback function on a remove 
- items count limiter
- use your hash function.

### Benchmark test  
go-cache [https://github.com/pmylund/go-cache](https://github.com/pmylund/go-cache)  
go version go1.3, single thread      
<pre>
Set
Cergoo.cache:    5000	    456282 ns/op   23200 B/op	    1450 allocs/op
go-cache:        5000	    734426 ns/op   66266 B/op	    2956 allocs/op
Get
Cergoo.cache:    5000	    359400 ns/op   23200 B/op	    1450 allocs/op
go-cache:        5000	    362525 ns/op   15184 B/op	     949 allocs/op
Inc
Cergoo.cache:    5000	    406278 ns/op   23200 B/op	    1450 allocs/op
go-cache:        5000	    428155 ns/op   15184 B/op	     949 allocs/op
</pre>

## chansubscriber
Subcribe channel pack, send messages of writer to a each subscribers.  
http://godoc.org/github.com/Cergoo/gol/chansubscriber

## counter 
Easy atomic counter type.    
http://godoc.org/github.com/Cergoo/gol/chansubscriber      

## err
Editable error implementation.  
http://godoc.org/github.com/Cergoo/gol/err

## fastbuf
io.Writer implementation.  
http://godoc.org/github.com/Cergoo/gol/fastbuf

###Comparition benchmark test
<pre>
Write
fastbuf:      10000000	       169 ns/op       0 B/op	       0 allocs/op
bytes.Buffer: 10000000	       222 ns/op       0 B/op	       0 allocs/op
</pre>

## filepath
Filepath util.  
http://godoc.org/github.com/Cergoo/gol/filepath
    
## hash
Hash functions library.  
http://godoc.org/github.com/Cergoo/gol/hash
  
## http/genid
Generate ID pkg  
  
    // NewHTTPGen ID creator, resize to base64 encoding, len(id) = 4*length/3.   
    // the actual length can be less per unit.      
    func NewHTTPGen(length uint8) HTTPGenID
 
    // Generate random strind http compatible.       
    func (t HTTPGenID) NewID() string

## http/methods
http methods name  
http://godoc.org/github.com/Cergoo/gol/http/methods

## http/cookie

    // Create new *http.Cookie  
    func NewCookie(name, value string, options *Options) *http.Cookie

    // Set cookie  
    func SetCookie(w http.ResponseWriter, name, value string, options *Options)

    // Del cookie   
    func DelCookie(w http.ResponseWriter, name string)
    
## http/session
Cookie based session engin implementation  

    // Constructor session engin  
    func NewSessionEngin(lenID uint8, ipProtect bool, stor TStor) *TSession 

    // Create new session  
    func (t *TSession) New(w http.ResponseWriter, r *http.Request, data interface{}) (id string)       

    // Delete session  
    func (t *TSession) Del(w http.ResponseWriter, r *http.Request)

    // Get session  
    func (t *TSession) Get(w http.ResponseWriter, r *http.Request) (id []byte, val interface{})

## http/router
Routing a path url to action. First elemet path is action name, others elemets is a request parameters

## jsonConfig
Support comments in json config files.      
http://godoc.org/github.com/Cergoo/gol/jsonConfig
         
## keyid
String key to uint id and uint id to string key association pack. No save thread.

    // Constructor
    func New() *Tkeyid
    
    // Set cortege (key, id)
    func (t *Tkeyid) Set(key string, id uint)
    
    // Delete from key and return id
    func (t *Tkeyid) DelFromKey(key string) (id uint, ok bool)
   
    // Delete from id and return key
    func (t *Tkeyid) DelFromId(id uint) (key string, ok bool)
   
    // Get id from key
    func (t *Tkeyid) GetId(key string) (id uint, ok bool)
   
    // Get key from id
    func (t *Tkeyid) GetKey(id uint) (key string, ok bool)

## refl
Additional reflection functions pack

    /* Universal caller of functions */  
    
    type (  
      FuncMap   map[string]reflect.Value  
      FuncSlice []reflect.Value
    )
    
    // Add to function map
    func (t FuncMap) Add(name string, f interface{})
    
    // Add to function slice, return element id
    func (t *FuncSlice) Add(f interface{}) int 
    
    // Call from map and return interface{}
    func (t FuncMap) Calli(name string, params ...interface{}) []interface{}
    
    // Call from slice and return interface{}
    func (t FuncSlice) Calli(id int, params ...interface{}) []interface{} 

    // Call function from a function map
    func (t FuncMap) Call(name string, params ...interface{}) []reflect.Value
    
    // Call function from a function slice
    func (t FuncSlice) Call(id int, params ...interface{}) []reflect.Value

    /* Other reflection functions */
          
    // A resize to slice all types. It panics if v's Kind is not slice.    
    func SliceResize(pointToSlice interface{}, newCap int)

    // Return true if keys map1 == keys map2. It panics if v's Kind is not map.
    func MapKeysEq(map1, map2 interface{}) bool
    
    // If "v" is struct copy fields to "m" map[string]interface{} and return true else return false.
    // If "unexported" true copy all fields. 
    func StructToMap(v interface{}, m map[string]interface{}, unexported bool, prefix string) bool
    
    // IsStruct returns true if the given variable is a struct or a pointer to struct.
    func IsStruct(v interface{}) bool
    
    // Return true if v is chan, func, interface, map, pointer, or slice and v is nil
    func IsNil(v interface{}) bool
    
    // Return true if v is nil or empty
    func IsEmpty(v interface{}) bool
    
## test  
Test helper functions is a simple assertion wrapper for Go's built in "testing" package   
fork jmervine/GoT  
http://godoc.org/github.com/Cergoo/gol/jsonConfig
  

## tplEngin\i18n
i18n pkg.

### Feature:
Load from .json format store language resource.  
Support tag: include context and plural. Example: `Field {{0}} must be filled {{1}} {{plural appel 1}}`.     
See tplEngin\i18n\exaple for more details.

    // Create language obj
    func New() Ti18n
    
    // Loade language resources  
    func Load(patch string, pluralAccess bool)

    // Create new replacer from language resources  
    func (t Ti18n) NewReplacer(langName string) *TReplacer

    // Get lang  
    func (t *TReplacer) Lang() string

    // Print. Get phrase from a map store.  
    func (t *TReplacer) P(key string, context ...interface{}) []byte
    
    // Get plural. Use if Load (pluralAccess==true)  
    func (t *TReplacer) Plural(key string, count float64) string
    
## tplEngin\i18n\plural
Plural form rules

## tplEngin\parser
Parser util from i18n & tpl pkg

## tplEngin\tplengin
Templare engin.
Attention! Work not complete.
