# gol (Go Library)
(c) 2013-2016 Cergoo   
under terms of MIT license

## encode/binary
###primitive
Package primitive it's a binary encode/decode primitive elementary implementation
http://godoc.org/github.com/Cergoo/gol/encode/binary/primitive

### encodebinary
Package encodebinary it's a binary encode/decode implementation to fast and easy serialize data.   
http://godoc.org/github.com/Cergoo/gol/encode/binary/encodebinary

### encodebinaryFast
Package encodebinaryFast it's a like protobuf encoder/decoder genarator from encodebinary protocol. See example.     
http://godoc.org/github.com/Cergoo/gol/encode/binary/encodebinaryFast

## encode/json
This is a set of packages to be marshaled from 'go' to .json. 
Standard unit of refusing to encode json from hash map keys which are not strings, 
This package encode them as arrays and hash map whose string keys as objects. 
Also realized rapid encoding json alternative github.com/pquerna/ffjson    
features:   
- coding without error;    
- supported json.Marshaler interface.
- tag json supported

<pre>
Encode
std_json:  			500000	      4525 ns/op     error json: unsupported type: map[int]str
gol_encodejson:  	200000	     13359 ns/op     844 B/op	      45 allocs/op
gol_encodejsonFast: 1000000	      2062 ns/op      88 B/op	       5 allocs/op
</pre>


### encodejson
http://godoc.org/github.com/Cergoo/gol/encode/json/encodejson

### encodejsonFast
http://godoc.org/github.com/Cergoo/gol/encode/json/encodejsonFast


## cache/cacheStr, cache/cacheUint
Package cache it's a in-memory key-value store based of the thread-safe hasmap implementation (key type string and key type uint64) similar to memcached that is suitable for applications running on a single machine.    
Automatic lifetime management of records can be enabled or disabled. LRU caches are affected by the problem leaching records in intensive add, i.e. the records  permanently pushed do not linger in the cach. This package does not implement the LRU. In this implementation the time life records indicated for all the generated cache, specified time value is the size of time interval during which a new record is guaranteed to live in the cache. Then have a record of lives at least one time interval maximum of two time interval + can be implemented for the "if it's read then it lives" if a record is requested that her life is prolonged for the next time interval.  
http://godoc.org/github.com/Cergoo/gol/cache/cacheUint    
http://godoc.org/github.com/Cergoo/gol/cache/cacheStr    

### Feature:
- thread-safe
- faste and high availability
- increment/decrement command supported
- save/load operation supported
- mechanism of managing the lifetime data: time expirations (set for the entire cache) and options "if read then life", support callback function on a remove 
- items count limiter
- use your hash function.
  
go-cache [https://github.com/pmylund/go-cache](https://github.com/pmylund/go-cache)  
go version go1.3, single thread      
<pre>
Set
Cergoo.cacheStr:    5000	    456282 ns/op   23200 B/op	    1450 allocs/op
go-cache:           5000	    734426 ns/op   66266 B/op	    2956 allocs/op
Get
Cergoo.cacheStr:    5000	    359400 ns/op   23200 B/op	    1450 allocs/op
go-cache:           5000	    362525 ns/op   15184 B/op	     949 allocs/op
Inc
Cergoo.cacheStr:    5000	    406278 ns/op   23200 B/op	    1450 allocs/op
go-cache:           5000	    428155 ns/op   15184 B/op	     949 allocs/op
</pre>

## chansubscriber
Subcribe channel pack, send messages of writer to a each subscribers.  
http://godoc.org/github.com/Cergoo/gol/chansubscriber

## counter 
Easy atomic counter type.    
http://godoc.org/github.com/Cergoo/gol/counter      

## crypto/enigma 
Package enigma its a crypto encripter/decripter, base64.URL safety from web, with periodically changes key   
http://godoc.org/github.com/Cergoo/gol/crypto/enigma 

## err
Editable error implementation.  
http://godoc.org/github.com/Cergoo/gol/err

## fastbuf
io.Writer implementation.  
http://godoc.org/github.com/Cergoo/gol/fastbuf

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
  
## http
### genid
Generate http compatible ID.
http://godoc.org/github.com/Cergoo/gol/http/genid  
  
### method
Http methods name.  
http://godoc.org/github.com/Cergoo/gol/http/method

### clientCache
http/1.1 client side cache control pkg.  
http://godoc.org/github.com/Cergoo/gol/http/clientCache

### cookie
Cookie pkg.  
http://godoc.org/github.com/Cergoo/gol/http/cookie/cookie
        
### cryptedcookie
CryptedCookie pkg.  
http://godoc.org/github.com/Cergoo/gol/http/cookie/cryptedcookie

### session
Cookie based session engin implementation.  
http://godoc.org/github.com/Cergoo/gol/http/session

### router
Routing a path url to action or file. First elemet path is action name, others elemets is a request parameters.  
http://godoc.org/github.com/Cergoo/gol/http/router

####Features:
- routing to file;
- suppart http method for REST routing;
- logging a errors action to stderr.

####Route example:
    pubic/1/en
    ------- ---- --
    actionName/:id/:lang
    and
    getfile/path/to/file
          
## jsonConfig
Support comments in json config files.      
http://godoc.org/github.com/Cergoo/gol/jsonConfig
         
## keyid
String key to uint id and uint id to string key association pack. No save thread.
http://godoc.org/github.com/Cergoo/gol/keyid

## reflect
### refl
Additional reflection functions  
http://godoc.org/github.com/Cergoo/gol/reflect/refl

### caller
Universal caller of functions  
http://godoc.org/github.com/Cergoo/gol/reflect/caller

### lookup
Package lookup it's a lookup reflection functions  
http://godoc.org/github.com/Cergoo/gol/reflect/lookup

## stack
### stack 
Package stack it's a simple stack implementation. No thread safe.    
http://godoc.org/github.com/Cergoo/gol/stack/stack
 
### stacklf
Package stacklf it's a implementation lockfree LIFO stack  
http://godoc.org/github.com/Cergoo/gol/stack/stacklf

### stackCounter
Package stackCounter it's a implementation lockfree LIFO stack under counter & limiter items    
http://godoc.org/github.com/Cergoo/gol/stack/stackCounter

### bytestack
Package bytestack it's a simple fixed length slice stack implementation. No thread safe.    
http://godoc.org/github.com/Cergoo/gol/stack/bytestack

## sync

### spinlock 
Package spinlock it's a simple spin lock implementation     
http://godoc.org/github.com/Cergoo/gol/sync/spinlock

### mrswUint 
Package mrswUint it's a multi reade single write dispatcher for uint key resource     
controller http://godoc.org/github.com/Cergoo/gol/sync/mrswUint/mrsw    
dispatche http://godoc.org/github.com/Cergoo/gol/sync/mrswUint/mrswd    

### mrswString 
Package mrswString it's a multi reade single write dispatcher for string key resource     
controller http://godoc.org/github.com/Cergoo/gol/sync/mrswString/mrsw    
dispatche http://godoc.org/github.com/Cergoo/gol/sync/mrswString/mrswd

## test  
Test helper functions is a simple assertion wrapper for Go's built in "testing" package, fork jmervine/GoT  
http://godoc.org/github.com/Cergoo/gol/jsonConfig
  
## tplEngin/i18n/i18n
Package i18n implementation  
http://godoc.org/github.com/Cergoo/gol/tplEngin/i18n/i18n   
Feature:  
    - Load from .json format language resource store.  
    - Support tag: include context vars and plural.   
    - Support user functions   
    - Support pluggable modules as a user functions librarys (a example 'humanize' mod implementation)       
Example: 
<pre>
Good afternoon, Mr.(Mrs.) {{0}}, you have {{1 %.2f}} {{plural apple 1}}.
Good afternoon, Mr.(Mrs.) {{0}}, you have {{f humanByteLong 1}}.
</pre>    
See tplEngin/i18n/i18n/exaple for more details. 

## tplEngin/i18n/plural
Plural form rules, fork plural github.com/vube/i18n  
http://godoc.org/github.com/Cergoo/gol/tplEngin/i18n/plural

## tplEngin/i18n/human
Package human it's a formatters for units to human friendly sizes  
http://godoc.org/github.com/Cergoo/gol/tplEngin/i18n/human

## tplEngin/parser
Parser util from i18n & tpl pkg  
http://godoc.org/github.com/Cergoo/gol/tplEngin/parser

## tplEngin\tplengin
Templare engin.
Attention! Work not complete.
