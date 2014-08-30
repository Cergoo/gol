package control1

import (
	. "../../registry"
	"fmt"
	"net/http"
)

func Run(w http.ResponseWriter, r *http.Request) {
	n := PoolResponseBody.Get().([]byte)
	n = append(n, []byte("action1 run")...)
	fmt.Fprint(w, string(n))
}
