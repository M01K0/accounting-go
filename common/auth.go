package common

import (
	"crypto/rsa"
	"github.com/alexyslozada/accounting-go/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strings"
	"errors"
	"github.com/alexyslozada/accounting-go/dao/executedao"
)

// Archivos para firmar y verificar los token
// openssl genrsa -out app.rsa 1024
// openssl rsa -in app.rsa -pubout > app.rsa.pub
const (
	privateKeyPath = "./keys/app.rsa"
	publicKeyPath  = "./keys/app.rsa.pub"
)

// Permiten firmar y verificar los token
var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func initKeys() {

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

// GenerateJWT genera un nuevo JWT token
func GenerateJWT(user models.User, method string) (string, error) {
	scopes, err := executedao.PermissionDAO.GetScopes(user.Profile.ID, method)
	if err != nil {
		return "", err
	}


	claims := models.AppClaims{
		User: user,
		Method: method,
		Scopes: scopes,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    "Contabilidad por Alexys",
		},
	}
	log.Printf("Creando un token a: %s, para el método: %s\n", user.Username, method)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	log.Println("Firmando el token...")
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

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
