package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetPublicRoutes(router *mux.Router) {
	router.Handle("/", http.FileServer(http.Dir("./public")))
}
