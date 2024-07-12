package models

type Team struct {
	ID    int
	Name  string
	Power int
	Pts   int
	W     int
	D     int
	L     int
	GD    int
}

type ChampionPrediction struct {
	TeamName       string  `json:"team_name"`
	WinProbability float64 `json:"win_probability"`
}