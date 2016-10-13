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
		data                                       LoginResource
		tokenPost, tokenPut, tokenDelete, tokenGet string
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

	// Se borra el password
	user.Passwd = ""

	// Se generan los token JWT
	tokenPost, err = common.GenerateJWT(user, "POST")
	if err != nil {
		common.DisplayError(w, err, "Error generando el access token", 500)
		return
	}
	tokenPut, err = common.GenerateJWT(user, "PUT")
	if err != nil {
		common.DisplayError(w, err, "Error generando el access token", 500)
		return
	}
	tokenDelete, err = common.GenerateJWT(user, "DELETE")
	if err != nil {
		common.DisplayError(w, err, "Error generando el access token", 500)
		return
	}
	tokenGet, err = common.GenerateJWT(user, "GET")
	if err != nil {
		common.DisplayError(w, err, "Error generando el access token", 500)
		return
	}

	authUser := models.AuthUser{
		User:        user,
		TokenPost:   tokenPost,
		TokenPut:    tokenPut,
		TokenDelete: tokenDelete,
		TokenGet:    tokenGet,
	}

	j, err := json.Marshal(AuthUserResource{authUser})
	if err != nil {
		common.DisplayError(w, err, "Error al generar la respuesta en json", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
