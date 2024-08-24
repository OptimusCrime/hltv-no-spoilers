package matches

import "time"

type Match struct {
	Url       string `json:"url"`
	Team1     string `json:"team1"`
	Team2     string `json:"team2"`
	EventName string `json:"eventName"`
	MatchType string `json:"type"`
}

type MatchGroup struct {
	MatchDate *time.Time `json:"date"`
	Matches   []Match    `json:"matches"`
}
