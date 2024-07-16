package models

import "time"

type Game struct {
	ID         int       `json:"id"`
	HomeTeamID int       `json:"homeTeamId"`
	AwayTeamID int       `json:"awayTeamId"`
	Date       time.Time `json:"date"`
	Location   string    `json:"location"`
}
