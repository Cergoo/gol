package control2

import (
	. "../../registry"
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter, r *http.Request) {
	n := PoolResponseBody.Get().([]byte)
	n = append(n, []byte("action2 run")...)
	fmt.Fprint(w, string(n))
}
