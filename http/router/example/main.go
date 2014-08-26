/*
	http://localhost:9999/action1/prm1/val1/prm2/val2
	http://localhost:9999/action1/prm1/val1/prm2
	http://localhost:9999/files/f1.txt
	http://localhost:9999/files
*/

// Example use pkg.
package main

import (
	"fmt"
	"github.com/Cergoo/gol/http/method"
	"github.com/Cergoo/gol/http/router"
	"log"
	"net/http"
	"time"
)

func action2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "action2", " ", r.Form.Encode(), " ", r.URL.Path)
}

func action1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "action1", " ", r.Form.Encode())
}

func main() {

	r := router.New()
	r.ServeFiles("files", "./directoryfiles")
	r.Handler(method.Get, "", action1)
	r.Handler(method.Get, "action1/id/lang", action1)
	r.Handler(method.Get, "action2/id/lang", action2)

	srv_htpp := &http.Server{
		Addr:           ":9999",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 18,
	}
	log.Fatal(srv_htpp.ListenAndServe())

}
