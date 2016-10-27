package common

import (
	"crypto/rsa"
	"github.com/alexyslozada/accounting-go/dao/executedao"
	"github.com/alexyslozada/accounting-go/models"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
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
		User:   user,
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
