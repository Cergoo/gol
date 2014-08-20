/*
	Routing a path url to action or file.
	(c) 2014 Cergoo
	under terms of ISC license
	==========================
	Description:
	First elemet path is action name, others elemets is a request parameters
	Features:
	- routing to file;
	- suppart http method for REST routing;
	- logging a errors action to stderr.
	Route example:
	pubic/1/en
	------- ---- --
	actionName/:id/:lang
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
	troute struct {
		prm    []string
		action func(http.ResponseWriter, *http.Request)
	}

	Trouter struct {
		routes            map[string]*troute
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
		routes:          make(map[string]*troute),
		filesRouteLabel: filesRouteLabel,
		errorLog:        log.New(os.Stderr, "router: ", log.LstdFlags),
	}
	if filesRootDir > "" {
		r.fileServerHandler = http.StripPrefix("/"+filesRouteLabel, http.FileServer(http.Dir(filesRootDir)))
	}

	return r
}

func (t *Trouter) AddRout(method, patch string, action func(w http.ResponseWriter, r *http.Request)) {
	parts := strings.Split(patch, "/")
	r := &troute{action: action}
	r.prm = append(r.prm, parts[1:]...)
	t.routes[method+parts[0]] = r
}

// Routing
func (t *Trouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(w, "error: %s\nat:\n%s", e, debug.Stack())
			t.errorLog.Printf("error: %s\nat:\n%s", e, debug.Stack())
		}
	}()
	url := r.URL.Path[1:]
	urlParts := strings.SplitN(url, "/", 30)

	// Files stream
	if t.fileServerHandler != nil && urlParts[0] == t.filesRouteLabel {
		t.fileServerHandler.ServeHTTP(w, r)
		return
	}

	// Find action
	action := t.routes[r.Method+urlParts[0]]
	if action == nil {
		fmt.Fprint(w, "url not found")
		return
	}

	// Parse url to paramrters
	r.ParseForm()
	urlParts = urlParts[1:]

	// find minimal len of a slice
	count := len(urlParts)
	if len(action.prm) < count {
		count = len(action.prm)
	}

	for id := 0; id < count; id++ {
		if len(urlParts[id]) > 0 {
			r.Form.Set(action.prm[id], urlParts[id])
		}
	}

	// Run action
	action.action(w, r)
}
