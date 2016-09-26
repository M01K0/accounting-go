package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Usuarios
	SetUserRoutes(router)

	return router
}
