/*
	url http routing to action
	(c) 2014 Cergoo
	under terms of ISC license
	==========================
	route example:
	getpage/lang/en
	------- ---- --
	actionName/paramName/paramValue
*/
package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

type (
	Trouter struct {
		Routes          map[string]func(http.ResponseWriter, *http.Request)
		filesRouteLabel string
		filesPathPrefix string
		errorLog        *log.Logger
	}
)

// constructor
func New(filesRouteLabel, filesPathPrefix string) *Trouter {
	return &Trouter{
		Routes:          make(map[string]func(http.ResponseWriter, *http.Request)),
		filesRouteLabel: filesRouteLabel,
		filesPathPrefix: filesPathPrefix,
		errorLog:        log.New(os.Stderr, "router: ", log.LstdFlags),
	}
}

// routing
func (t *Trouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(w, "error: %s\nat:\n%s", e, debug.Stack())
			t.errorLog.Printf("error: %s\nat:\n%s", e, debug.Stack())
		}
	}()

	url := r.URL.Path[1:]
	urlParts := strings.SplitN(url, "/", 30)

	// files stream
	if urlParts[0] == t.filesRouteLabel {
		http.ServeFile(w, r, t.filesPathPrefix+strings.TrimPrefix(url, t.filesRouteLabel))
		return
	}

	// find action
	action := t.Routes[urlParts[0]]
	if action == nil {
		fmt.Fprint(w, "url not found")
		return
	}

	// parse url paramrters
	r.ParseForm()
	for id := 1; id < len(urlParts); id = +2 {
		r.Form.Add(urlParts[id], urlParts[id+1])
	}

	// run action
	action(w, r)
	return
}
