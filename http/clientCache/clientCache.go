/*
	http/1.1 client side cache control pkg
	(c) 2014 Cergoo
	under terms of ISC license
*/

package clientCache

import (
	"net/http"
	"strconv"
	"time"
)

// return true if cache validate (304)
func Cache(w http.ResponseWriter, rh http.Header, timelive int, noStore bool, cacheControl, etag string) bool {
	header := w.Header()
	header.Set("Cache-Control", "max-age="+strconv.Itoa(timelive)+", "+cacheControl)
	header.Set("Expires", time.Now().Add(time.Duration(timelive)*time.Second).Format(http.TimeFormat))

	if !noStore {
		header.Set("ETag", etag)
		if etag == rh.Get("If-None-Match") {
			w.WriteHeader(http.StatusNotModified)
			return true
		}
	}

	return false
}
