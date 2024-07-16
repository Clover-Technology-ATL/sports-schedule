package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Clover-Technology-ATL/sports-schedule/backend/models"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	// Simulated data (replace with actual data from a database)
	players := []models.Player{
		{ID: 1, Name: "Player 1", TeamID: 1},
		{ID: 2, Name: "Player 2", TeamID: 2},
		{ID: 3, Name: "Player 3", TeamID: 1},
	}

	json.NewEncoder(w).Encode(players)
}
