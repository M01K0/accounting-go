package routers

import (
	"github.com/alexyslozada/accounting-go/controllers"
	"github.com/gorilla/mux"
)

func SetLoginRoutes(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
