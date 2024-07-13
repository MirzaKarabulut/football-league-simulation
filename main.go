package main

import (
	"football-league-simulation/db"
	"football-league-simulation/routes"
	"log"
	"net/http"
)

// @title Insider API
// @version 1.0
// @description This is a simple football league simulation API.
// @host localhost:8080
// @BasePath /

func main() {
    db.Init()
    router := routes.InitRoutes()

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
