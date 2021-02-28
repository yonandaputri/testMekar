package config

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testMekar/tools"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router)  {
	host := tools.GetEnv("API_HOST", "localhost")
	port := tools.GetEnv("API_PORT", "8080")

	fmt.Println("Starting Web Server at " + host + ":" + port)
	err := http.ListenAndServe(host+ ":" + port, router)
	if err != nil {
		log.Fatal(err)
	}
}