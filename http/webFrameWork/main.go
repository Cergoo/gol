/*
	http://localhost:9999/action1/prm1/val1/prm2/val2
	http://localhost:9999/action1/prm1/val1/prm2
	http://localhost:9999/files/f1.txt
	http://localhost:9999/files
*/

// Example use pkg.
package main

import (
	"./controllers/control2"
	"github.com/Cergoo/gol/http/method"
	"github.com/Cergoo/gol/http/router"
	"github.com/Cergoo/gol/http/webFrameWork/controllers/control1"
	"log"
	"net/http"
	"time"
)

func main() {

	r := router.New(nil)
	r.FileServer("files", "./directoryfiles")
	r.Handler(method.Get, "", control1.Run)
	r.Handler(method.Get, "action1/id/lang", control1.Run)
	r.Handler(method.Get, "action2/id/lang", control2.Run)

	srv_htpp := &http.Server{
		Addr:           ":9999",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 18,
	}
	log.Fatal(srv_htpp.ListenAndServe())

}
