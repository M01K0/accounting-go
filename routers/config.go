package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Usuarios
	SetLoginRoutes(router)
	// Centros de Costo
	SetCostCenterRoutes(router)

	return router
}
