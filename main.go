package main

import (
	"football-league-simulation/db"
	"football-league-simulation/routes"
	"log"
	"net/http"
)

func main() {
    db.Init()
    router := routes.InitRoutes()

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
