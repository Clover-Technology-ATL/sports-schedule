package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Clover-Technology-ATL/sports-schedule/backend/models"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	// Simulated data (replace with actual data from a database)
	teams := []models.Team{
		{ID: 1, Name: "Team 1"},
		{ID: 2, Name: "Team 2"},
		{ID: 3, Name: "Team 3"},
	}

	json.NewEncoder(w).Encode(teams)
}
