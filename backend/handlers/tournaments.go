package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Clover-Technology-ATL/sports-schedule/backend/models"
)

func GetTournaments(w http.ResponseWriter, r *http.Request) {
	// Simulated data (replace with actual data from a database)
	tournaments := []models.Tournament{
		{ID: 1, Name: "Tournament 1"},
		{ID: 2, Name: "Tournament 2"},
	}

	json.NewEncoder(w).Encode(tournaments)
}
