package routers

import (
	"github.com/gorilla/mux"
	"github.com/alexyslozada/accounting-go/controllers"
	"github.com/urfave/negroni"
	"github.com/alexyslozada/accounting-go/common"
)

func SetCostCenterRoutes(router *mux.Router) {
	prefix := "/api/cost-centers"
	ccr := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	ccr.HandleFunc("/", controllers.CostCenterCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.HandlerFunc(common.MakeMiddleWarePermissions(prefix, common.IsPermitted)),
		negroni.Wrap(ccr),
	))
}
