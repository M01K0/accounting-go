package main

import (
	"github.com/alexyslozada/accounting-go/common"
	"github.com/alexyslozada/accounting-go/routers"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	// Inicia la lógica
	common.StartUp()

	// Inicia los router
	router := routers.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Iniciado...")
	server.ListenAndServe()
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
