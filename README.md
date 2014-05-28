# gol (Go Library)
(c) 2013-2014 Cergoo   
under terms of ISC license


## cache
Cache is an in-memory key:value store/cache similar to memcached that is suitable for applications running on a single machine.    
The package is the implementation of hashmap for organizing in-memory key-value data store. Automatic lifetime management of records can be enabled or disabled. LRU caches are affected by the problem leaching records in intensive add, i.e. the records  permanently pushed do not linger in the cach. This package does not implement the LRU. In this implementation the time life records indicated for all the generated cache, specified time value is the size of time interval during which a new record is guaranteed to live in the cache. Then have a record of lives at least one time interval maximum of two time interval + can be implemented for the "if it's read then it lives" if a record is requested that her life is prolonged for the next time interval.  
    
### Feature:
- thread-safe
- faste and high availability
- increment/decrement command supported
- save/load operation supported
- mechanism of managing the lifetime data: time expirations (set for the entire cache) and options "if read then life"
- support callback function on a timer remove automatically
- items count limiter
- use your hash function.

### Comparition benchmark test  
go-cache [https://github.com/pmylund/go-cache](https://github.com/pmylund/go-cache)  
go version go1.2.1, single thread  
<pre>
Set
Cergoo.cache:    5000	    646846 ns/op   24000 B/op	    3000 allocs/op
go-cache:        2000	   1101513 ns/op   66227 B/op	    4011 allocs/op
Get
Cergoo.cache:    5000	    546850 ns/op   24000 B/op	    3000 allocs/op
go-cache:        5000	    546850 ns/op   16000 B/op	    2000 allocs/op
Inc
Cergoo.cache:    5000	    612472 ns/op   24000 B/op	    3000 allocs/op
go-cache:        5000	    590598 ns/op   16000 B/op	    2000 allocs/op
</pre>

## cookie

    // Create new *http.Cookie  
    func NewCookie(name, value string, options *Options) *http.Cookie

    // Set cookie  
    func SetCookie(w http.ResponseWriter, name, value string, options *Options)

    // Del cookie   
    func DelCookie(w http.ResponseWriter, name string)

## counter 
Easy atomic counter type  

    // Get current count value  
    func (t *T_counter) Get() uint64  

    // Set current count value  
    func (t *T_counter) Set(v uint64)  

    // Increment  
    func (t *T_counter) Inc() uint64  

    // Decrement  
    func (t *T_counter) Dec() uint64  

    // Add value    
    func (t *T_counter) Add(v uint64, dec bool) uint64  

    // Get current limit value  
    func (t *T_counter) GetLimit() uint64  

    // Set new limit value  
    func (t *T_counter) SetLimit(v uint64) 

    // Check limit value  
    func (t *T_counter) Check() bool 

    // Check limit value  
    func (t *T_counter) Check1(v uint64) bool  

## err
Editable error implementation

## fastbuf
io.Writer implementation  

    // write to buf
    func (t *Buf) Write(p []byte) (n int, err error)
    
    // get all buf and clear buf
    func (t *Buf) Flush() (r []byte)  

###Comparition benchmark test
<pre>
Write
fastbuf:      5000000        550 ns/op       0 B/op	       0 allocs/op
bytes.Buffer: 1000000       1099 ns/op       0 B/op	       0 allocs/op
</pre>

## filepath
Filepath util
    
    // modified function Ext standart "path/filepath" pkg  
    func Ext(fullname string) (name, ext string)

## genid
Generate ID pkg  
  
    // NewHTTPGen ID creator, resize to base64 encoding, len(id) = 4*length/3.   
    // the actual length can be less per unit.      
    func NewHTTPGen(length uint16) HTTPGenID
 
    // Generate random strind http compatible.       
    func (t HTTPGenID) NewID() string

## hash
Hash functions library

    // FAQ6 hash  
    func HashFAQ6(str []byte) (h uint32)  

    // Rot13 hash  
    func HashRot13(str []byte) (h uint32)  

    // Ly hash  
    func HashLy(str []byte) (h uint32)  
  
    // Rs hash  
    func HashRs(str []byte) (h uint32)  

## jsonConfig
Support comments in json config files.    

    // Load & remove comments from source .json file  
    func Load(fromPath string, toVar interface{})

    // Remove comments from source .json  
    func RemoveComment(source []byte) (result []byte)     

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
    
    // If "v" is struct copy fields to map[string]interface{} "m" and return true else return false
    func StructToMap(v interface{}, m map[string]interface{}, prefix string) bool

## session
Cookie based session engin implementation  

    // Constructor session engin  
    func NewSessionEngin(lenID uint8, stor TStor) *TSession 

    // Create new session  
    func (t *TSession) New(w http.ResponseWriter, data interface{}) (id string)  

    // Delete session  
    func (t *TSession) Del(w http.ResponseWriter, r *http.Request)

    // Get session  
    func (t *TSession) Get(w http.ResponseWriter, r *http.Request) interface{}

## test  
Test helper functions is a simple assertion wrapper for Go's built in "testing" package   
fork jmervine/GoT

    // constructor
    func New(t *testing.T) *TT

    // Equivalence check   
    func (t *TT) Eq(id string, a, b interface{})
    
    // No Equivalent check
    func (t *TT) NoEq(id string, a, b interface{})  

## tplEngin\i18n
i18n pkg.

### Feature:
Load from .json format store language resource.  
Support tag: include context and plural. Example: `Field {{0}} must be filled {{1}} {{plural appel 1}}`.   
Support map (type key string) and slice (type key int) access to phrase.  
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