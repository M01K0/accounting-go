package common

import (
	"errors"
	"github.com/alexyslozada/accounting-go/models"
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

		scopes, ok := s.([]models.Scope)
		method := r.Method
		var result bool

		for _, p := range scopes {
			if p.Path == path {
				for _, m := range p.Methods {
					if m == method {
						result = true
						break
					}
				}
				break
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
