package common

import (
	"errors"
	"github.com/alexyslozada/accounting-go/dao/executedao"
	"github.com/alexyslozada/accounting-go/models"
	"github.com/gorilla/context"
	"net/http"
)

func IsPermitted(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func MakeMiddleWarePermissions(path string, fn func(http.ResponseWriter, *http.Request, http.HandlerFunc)) (func(http.ResponseWriter, *http.Request, http.HandlerFunc)) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		u, ok := context.GetOk(r, "user")
		if !ok {
			DisplayError(w, errors.New("No se encuentra el usuario en el token"), "No se encuentra el usuario en el token", http.StatusBadRequest)
			return
		}

		us, ok := u.(models.User)
		if !ok {
			DisplayError(w, errors.New("El usuario no tiene la estructura correcta"), "El usuario no tiene la estructura correcta", http.StatusBadRequest)
			return
		}

		method := r.Method
		result, err := executedao.PermissionDAO.IsPermitted(us.Profile, path, method)
		if err != nil {
			DisplayError(w, err, "Error al buscar los permisos del sistema", http.StatusInternalServerError)
			return
		}
		if result {
			fn(w, r, next)
		} else {
			DisplayError(w, errors.New("No estás autorizado"), "No estás autorizado para esta opción", http.StatusUnauthorized)
			return
		}
	}
}
