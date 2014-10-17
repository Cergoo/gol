package control1

import (
	"fmt"
	. "github.com/Cergoo/gol/http/webFrameWork/registry"
	"net/http"
)

func Run(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	handler.
	n := PoolResponseBody.Get().([]byte)
	n = append(n, []byte("action1 run")...)
	fmt.Fprint(w, string(n))
}
