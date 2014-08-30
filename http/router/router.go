// (c) 2014 Cergoo
// under terms of ISC license

/*
Package router it's http routing a path url to action or file implementation.
Description: first elemet path is action name, others elemets is a request parameters.
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
	"path"
	"runtime/debug"
	"strings"
)

//
const (
	actionTypeAction = iota
	actionTypeFile
)

type (
	troute struct {
		prm        []string
		action     http.HandlerFunc
		actionType uint8
	}

	// Trouter router
	Trouter struct {
		routes   map[string]*troute
		errorLog *log.Logger
		notFound http.HandlerFunc
	}
)

// New constructor new router
func New(notFound http.HandlerFunc) *Trouter {
	if notFound == nil {
		notFound = http.NotFound
	}
	return &Trouter{
		routes:   make(map[string]*troute),
		errorLog: log.New(os.Stderr, "router: ", log.LstdFlags),
		notFound: notFound,
	}
}

// FileServer set serve files
func (t *Trouter) FileServer(label, root string) {
	label = strings.ToLower(label)
	t.routes[method.Get+label] = &troute{
		action:     fileServer(label, root),
		actionType: actionTypeFile,
	}
}

func fileServer(prefix, root string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, root+path.Clean(r.URL.Path[len(prefix)+1:]))
	}
}

// Handler set hadler
func (t *Trouter) Handler(method, patch string, action func(w http.ResponseWriter, r *http.Request)) {
	r := &troute{
		action:     action,
		actionType: actionTypeAction,
	}
	parts := strings.Split(patch, "/")
	r.prm = append(r.prm, parts[1:]...)
	t.routes[method+strings.ToLower(parts[0])] = r
}

// ServeHTTP routing function
func (t *Trouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(w, "error: %s\nat:\n%s", e, debug.Stack())
			t.errorLog.Printf("error: %s\nat:\n%s", e, debug.Stack())
		}
	}()

	urlParts := strings.SplitN(r.URL.Path[1:], "/", 50)

	// Find action
	action := t.routes[r.Method+strings.ToLower(urlParts[0])]
	// Action nil
	if action == nil {
		t.notFound(w, r)
		return
	}
	// Action get file
	if action.actionType == actionTypeFile {
		action.action(w, r)
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
	action.action(w, r)
}
