package controllers

import (
	"encoding/json"
	"github.com/alexyslozada/accounting-go/common"
	"github.com/alexyslozada/accounting-go/dao/executedao"
	"github.com/alexyslozada/accounting-go/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		data  LoginResource
		token string
	)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.DisplayError(w, err, "Estructura de login no válida.", 401)
		return
	}

	login := data.Data
	user := models.User{Email: login.Email, Passwd: login.Password}
	err = executedao.LoginDAO.Login(&user)
	if err != nil {
		common.DisplayError(w, err, "Usuario/Clave inválidos o inactivo", 401)
		return
	}

	user.Passwd = "" // Se borra el password
	// Se genera el token JWT
	token, err = common.GenerateJWT(user)
	if err != nil {
		common.DisplayError(w, err, "Error generando el access token", 500)
		return
	}

	authUser := models.AuthUser{user, token}
	j, err := json.Marshal(AuthUserResource{authUser})
	if err != nil {
		common.DisplayError(w, err, "Error al generar la respuesta en json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
