package common

import (
	"net/http"
)

// QueryValues devuelve el valor o valores
// de una petición en la URL via Query (?)
func QueryValues(req *http.Request, param string) []string {
	vals := req.URL.Query()
	val := vals[param]
	return val
}

