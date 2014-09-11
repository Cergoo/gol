# gol (Go Library)
(c) 2013-2014 Cergoo   
under terms of ISC license

## binaryED/primitive
Package primitive it's a binary encode/decode primitive elementary implementation
http://godoc.org/github.com/Cergoo/gol/binaryED/primitive

## binaryED/binaryED
Package binaryED it's a binary encode/decode implementation to serialize data.   
(fork github.com/youtube/vitess/go/bson)    
http://godoc.org/github.com/Cergoo/gol/binaryED/binaryED

## cache
Package cache it's a in-memory key-value store based of the thread-safe hasmap implementation similar to memcached that is suitable for applications running on a single machine.    
Automatic lifetime management of records can be enabled or disabled. LRU caches are affected by the problem leaching records in intensive add, i.e. the records  permanently pushed do not linger in the cach. This package does not implement the LRU. In this implementation the time life records indicated for all the generated cache, specified time value is the size of time interval during which a new record is guaranteed to live in the cache. Then have a record of lives at least one time interval maximum of two time interval + can be implemented for the "if it's read then it lives" if a record is requested that her life is prolonged for the next time interval.  
http://godoc.org/github.com/Cergoo/gol/cache

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
http://godoc.org/github.com/Cergoo/gol/counter      

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
Generate http compatible ID.
http://godoc.org/github.com/Cergoo/gol/http/genid  
  
## http/method
Http methods name.  
http://godoc.org/github.com/Cergoo/gol/http/method

## http/clientCache
http/1.1 client side cache control pkg.  
http://godoc.org/github.com/Cergoo/gol/http/clientCache

## http/cookie
Cookie pkg.  
http://godoc.org/github.com/Cergoo/gol/http/cookie
        
## http/session
Cookie based session engin implementation.  
http://godoc.org/github.com/Cergoo/gol/http/session

## http/router
Routing a path url to action or file. First elemet path is action name, others elemets is a request parameters.  
http://godoc.org/github.com/Cergoo/gol/http/router

###Features:
- routing to file;
- suppart http method for REST routing;
- logging a errors action to stderr.

###Route example:
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

## lockfree/stack
Package stack it's a implementation lockfree LIFO stack  
http://godoc.org/github.com/Cergoo/gol/lockfree/stack

## reflect/refl
Additional reflection functions  
http://godoc.org/github.com/Cergoo/gol/reflect/refl

## reflect/caller
Universal caller of functions  
http://godoc.org/github.com/Cergoo/gol/reflect/caller

## reflect/lookup
Package lookup it's a lookup reflection functions  
http://godoc.org/github.com/Cergoo/gol/reflect/lookup

## stack 
Package stack it's a implementation lockfree LIFO stack under counter & limiter items    
http://godoc.org/github.com/Cergoo/gol/stack

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
