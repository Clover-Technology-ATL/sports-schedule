package main

import (
	"log"
	"net/http"

	"github.com/Clover-Technology-ATL/sports-schedule/backend/handlers"
)

func main() {
	http.HandleFunc("/api/teams", handlers.GetTeams)
	http.HandleFunc("/api/players", handlers.GetPlayers)
	http.HandleFunc("/api/games", handlers.GetGames)
	http.HandleFunc("/api/tournaments", handlers.GetTournaments)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
