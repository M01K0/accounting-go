package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Archivos estáticos
	SetPublicRoutes(router)
	SetFaviconRoute(router)
	// Usuarios
	SetLoginRoutes(router)
	// Centros de Costo
	SetCostCenterRoutes(router)

	return router
}
