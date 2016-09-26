package main

import (
	"log"
	"github.com/urfave/negroni"
	"github.com/alexyslozada/accounting-go/common"
	"github.com/alexyslozada/accounting-go/routers"
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
		Addr: ":8080",
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
