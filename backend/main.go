package main

import (
	"testMekar/config"
	"testMekar/master"
)

func main() {
	db, _ := config.ConnectionDB()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}
