// (c) 2014 Cergoo
// under terms of ISC license

/*
Routing a path url to action or file.
Description: First elemet path is action name, others elemets is a request parameters
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
	"github.com/Cergoo/gol/http/method"
	"github.com/Cergoo/gol/util"
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
		routes   map[string]*troute
		errorLog *log.Logger
	}
)

// Construcor
func New() *Trouter {
	return &Trouter{
		routes:   make(map[string]*troute),
		errorLog: log.New(os.Stderr, "router: ", log.LstdFlags),
	}
}

// Set serve files
func (t *Trouter) ServeFiles(label, root string) {
	t.routes[method.Get+label] = &troute{action: http.StripPrefix("/"+label, http.FileServer(http.Dir(root))).ServeHTTP}
}

// Set hadler
func (t *Trouter) Handler(method, patch string, action func(w http.ResponseWriter, r *http.Request)) {
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

	urlParts := strings.SplitN(r.URL.Path[1:], "/", 20)

	// Find action
	action := t.routes[r.Method+urlParts[0]]
	if action == nil {
		http.NotFound(w, r)
		return
	}

	// Parse url to parameters
	r.ParseForm()
	urlParts = urlParts[1:]
	count := util.Min(len(urlParts), len(action.prm))
	for id := 0; id < count; id++ {
		if len(urlParts[id]) > 0 {
			r.Form.Set(action.prm[id], urlParts[id])
		}
	}

	// Run action
	action.action(w, r)
}
