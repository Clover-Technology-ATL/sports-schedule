package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Clover-Technology-ATL/sports-schedule/backend/models"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	// Simulated data (replace with actual data from a database)
	games := []models.Game{
		{ID: 1, HomeTeamID: 1, AwayTeamID: 2, Date: time.Now(), Location: "Stadium A"},
		{ID: 2, HomeTeamID: 3, AwayTeamID: 1, Date: time.Now().AddDate(0, 0, 7), Location: "Stadium B"},
	}

	json.NewEncoder(w).Encode(games)
}
