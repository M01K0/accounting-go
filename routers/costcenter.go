package routers

import (
	"github.com/alexyslozada/accounting-go/common"
	"github.com/alexyslozada/accounting-go/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetCostCenterRoutes(router *mux.Router) {
	prefix := "/api/cost-centers"
	ccr := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	ccr.HandleFunc("/", controllers.CostCenterInsert).Methods("POST")
	ccr.HandleFunc("/", controllers.CostCenterGetAll).Methods("GET")
	ccr.HandleFunc("/{id:[0-9]+}", controllers.CostCenterUpdate).Methods("PUT")
	ccr.HandleFunc("/{id:[0-9]+}", controllers.CostCenterGetByID).Methods("GET")
	ccr.HandleFunc("/{id:[0-9]+}", controllers.CostCenterDelete).Methods("DELETE")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.HandlerFunc(common.MakeMiddleWarePermissions(prefix, common.IsPermitted)),
		negroni.Wrap(ccr),
	))
}
