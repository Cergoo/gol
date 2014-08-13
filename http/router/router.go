/*
	url http routing to action
	(c) 2014 Cergoo
	under terms of ISC license
	==========================
	route example:
	getpage/lang/id
	------- ---- --
	routeID param 1,2...
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
		routes          map[string]*troute
		filesRouteLabel string
		filesPathPrefix string
		errorLog        *log.Logger
	}

	troute struct {
		action func(http.ResponseWriter, *http.Request)
		param  []string // url parts to request parameters
	}
)

// constructor
func New(filesRouteLabel, filesPathPrefix string) *Trouter {
	return &Trouter{
		routes:          make(map[string]*troute),
		filesRouteLabel: filesRouteLabel,
		filesPathPrefix: filesPathPrefix,
		errorLog:        log.New(os.Stderr, "router: ", log.LstdFlags),
	}
}

/*
	Add route
	getpage/lang/id
	------- ---- --
	routeID param 1,2
*/
func (t *Trouter) RouteAdd(route string, action func(http.ResponseWriter, *http.Request)) {
	routeParts := strings.Split(route, "/")
	var param []string
	if len(routeParts) > 1 {
		param := make([]string, len(routeParts)-1)
		copy(param, routeParts[1:])
	}
	t.routes[routeParts[0]] = &troute{action: action, param: param}
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

	route := t.routes[urlParts[0]]
	if route == nil {
		fmt.Fprint(w, "url not found")
		return
	}
	urlParts = urlParts[1:]
	if len(route.param) > len(urlParts) {
		fmt.Fprint(w, "url not valid")
		return
	}
	r.ParseForm()
	for id, val := range route.param {
		r.Form.Add(val, urlParts[id])
	}

	route.action(w, r)
	return
}
