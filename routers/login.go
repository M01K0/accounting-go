package routers

import (
	"github.com/gorilla/mux"
	"github.com/alexyslozada/accounting-go/controllers"
)

func SetLoginRoutes(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
