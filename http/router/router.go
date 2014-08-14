/*
	Routing a path url to action or file.
	(c) 2014 Cergoo
	under terms of ISC license
	==========================
	Description:
	First elemet path is action name, others elemets is request parameters of a type: name/value
	Features:
	- routing to file;
	- routing to action;
	- logging a errors action to stderr.
	Route example:
	getpage/lang/en
	------- ---- --
	actionName/paramName/paramValue
	and
	getfile/path/to/file
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
		Routes            map[string]func(http.ResponseWriter, *http.Request)
		filesRouteLabel   string
		fileServerHandler http.Handler
		errorLog          *log.Logger
	}
)

/*
	constructor
	if filesRootDir == "" then no http files accesse
*/
func New(filesRouteLabel, filesRootDir string) *Trouter {
	r := &Trouter{
		Routes:          make(map[string]func(http.ResponseWriter, *http.Request)),
		filesRouteLabel: filesRouteLabel,
		errorLog:        log.New(os.Stderr, "router: ", log.LstdFlags),
	}
	if filesRootDir > "" {
		r.fileServerHandler = http.StripPrefix("/"+filesRouteLabel, http.FileServer(http.Dir(filesRootDir)))
	}

	return r
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
	if t.fileServerHandler != nil && urlParts[0] == t.filesRouteLabel {
		t.fileServerHandler.ServeHTTP(w, r)
		return
	}

	// find action
	action := t.Routes[urlParts[0]]
	if action == nil {
		fmt.Fprint(w, "url not found")
		return
	}

	// parse url to paramrters
	r.ParseForm()
	count := len(urlParts) - 1
	for id := 1; id < count; id += 2 {
		r.Form.Add(urlParts[id], urlParts[id+1])
	}

	// run action
	action(w, r)
}
