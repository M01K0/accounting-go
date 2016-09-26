package routers

import (
	"github.com/gorilla/mux"
	"github.com/alexyslozada/accounting-go/controllers"
)

func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
}
