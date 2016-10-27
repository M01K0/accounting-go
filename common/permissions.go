package common

import (
	"errors"
	"net/http"
	"strings"
	"github.com/alexyslozada/accounting-go/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go/request"
)

// Authorize Middleware para validar los JWT token
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Get token from request
	// El extractor podría ser: request.AuthorizationHeaderExtractor o tal vez el personalizado TokenFromAuthHeader
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Como solo tenemos una llave pública, la devolvemos
		return verifyKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayError(w, err, "Su token ha expirado, por favor vuelva a ingresar", http.StatusUnauthorized)
				return
			default:
				DisplayError(w, err, "Error en el token de acceso.", http.StatusUnauthorized)
				return
			}
		default:
			DisplayError(w, err, "Error al procesar el token.", http.StatusUnauthorized)
			return
		}
	}

	if token.Valid {
		context.Set(r, "user", token.Claims.(*models.AppClaims).User)
		context.Set(r, "methodrequest", token.Claims.(*models.AppClaims).Method)
		context.Set(r, "scopes", token.Claims.(*models.AppClaims).Scopes)
		next(w, r)
	} else {
		DisplayError(w, err, "Token de acceso inválido.", http.StatusUnauthorized)
	}
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}

func IsPermitted(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func MakeMiddleWarePermissions(path string, fn func(http.ResponseWriter, *http.Request, http.HandlerFunc)) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
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
