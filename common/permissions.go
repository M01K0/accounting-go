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

		m, ok := context.GetOk(r, "methodrequest")
		if !ok {
			DisplayError(w, errors.New("No se encuentra el método en el token"), "No se encuentra el método en el token", http.StatusBadRequest)
			return
		}
		methodRequest, ok := m.(string)
		if !ok {
			DisplayError(w, errors.New("Error al traer el tipo de método del token"), "Error al traer el tipo de método del token", http.StatusInternalServerError)
			return
		}
		method := r.Method
		if methodRequest != method {
			DisplayError(w, errors.New("El método del token no es el solicitado."), "El método del token no es el solicitado.", http.StatusBadRequest)
			return
		}

		s, ok := context.GetOk(r, "scopes")
		if !ok {
			DisplayError(w, errors.New("No se encuentran los permisos en el token"), "No se encuentra el usuario en el token", http.StatusBadRequest)
			return
		}

		scopes, ok := s.([]string)

		var result bool

		// Se busca en el slice la ruta sin el /api/
		for i := 0; i < len(scopes); i++ {
			if path[5:] == scopes[i] {
				result = true
				break
			}
		}

		if result {
			fn(w, r, next)
		} else {
			DisplayError(w, errors.New("No estás autorizado"), "No estás autorizado para esta opción", http.StatusForbidden)
			return
		}
	}
}
