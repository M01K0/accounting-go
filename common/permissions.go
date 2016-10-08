package common

import (
	"errors"
	"github.com/gorilla/context"
	"net/http"
)

func IsPermitted(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func MakeMiddleWarePermissions(path string, fn func(http.ResponseWriter, *http.Request, http.HandlerFunc)) (func(http.ResponseWriter, *http.Request, http.HandlerFunc)) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		s, ok := context.GetOk(r, "scopes")
		if !ok {
			DisplayError(w, errors.New("No se encuentran los permisos en el token"), "No se encuentra el usuario en el token", http.StatusBadRequest)
			return
		}

		scopes, ok := s.(map[string][]string)
		method := r.Method
		var result bool

		// Se busca en el map la ruta sin el /api
		if methods, ok := scopes[path[4:]]; ok {
			for _, m := range methods {
				if m == method {
					result = true
					break
				}
			}
		}

		if result {
			fn(w, r, next)
		} else {
			DisplayError(w, errors.New("No estás autorizado"), "No estás autorizado para esta opción", http.StatusUnauthorized)
			return
		}
	}
}
