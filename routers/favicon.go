package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/imgs/favicon.ico")
}

func SetFaviconRoute(router *mux.Router) {
	prefix := "/favicon.ico"
	r := mux.NewRouter()
	r.HandleFunc(prefix, faviconHandler).Methods("GET")

	router.PathPrefix(prefix).Handler(negroni.New(negroni.Wrap(r)))
}
