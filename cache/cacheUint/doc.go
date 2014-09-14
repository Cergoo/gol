// (c) 2013 Cergoo
// under terms of ISC license

/*
Package cacheUint it's a in-memory key-value store based of the thread-safe hasmap (key type Uint64) implementation similar to memcached that is suitable for applications running on a single machine.
Automatic lifetime management of records can be enabled or disabled. LRU caches are affected by the problem leaching records in intensive add, i.e. the records  permanently pushed do not linger in the cach. This package does not implement the LRU. In this implementation the time life records indicated for all the generated cache, specified time value is the size of time interval during which a new record is guaranteed to live in the cache. Then have a record of lives at least one time interval maximum of two time interval + can be implemented for the "if it's read then it lives" if a record is requested that her life is prolonged for the next time interval.
*/
package cacheUint
