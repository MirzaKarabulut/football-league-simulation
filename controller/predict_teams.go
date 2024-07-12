// controllers/controller.go

package controllers

import (
	"fmt"
	"football-league-simulation/db"
	"football-league-simulation/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func PredictChampionBeforeLastWeek(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	currentWeek := 4

	teams := fetchTeams()

	predictions := predictChampionProbabilities(teams)

	fmt.Printf("Before the last week (Week %d), predicted champions and probabilities:\n", currentWeek)
	for _, prediction := range predictions {
		fmt.Printf("Team: %s\n", prediction.TeamName)
		fmt.Printf("Win Probability: %.2f%%\n", prediction.WinProbability)
	}
}

func fetchTeams() []*models.Team {
	rows, err := db.DB.Query("SELECT id, name, pts, w, d, l, gd FROM teams")
	if err != nil {
		fmt.Println("Error fetching teams:", err)
		return nil
	}
	defer rows.Close()

	var teams []*models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.ID, &team.Name, &team.Pts, &team.W, &team.D, &team.L, &team.GD)
		if err != nil {
			fmt.Println("Error scanning team row:", err)
			continue
		}
		teams = append(teams, &team)
	}

	return teams
}

func predictChampionProbabilities(teams []*models.Team) []*models.ChampionPrediction {
	if len(teams) == 0 {
		fmt.Println("No teams found to predict champion.")
		return nil
	}

	totalPts := 0
	for _, team := range teams {
		totalPts += team.Pts
	}

	var predictions []*models.ChampionPrediction
	for _, team := range teams {
		winProbability := float64(team.Pts) / float64(totalPts) * 100.0
		predictions = append(predictions, &models.ChampionPrediction{
			TeamName:       team.Name,
			WinProbability: winProbability,
		})
	}


	sumProbabilities := 0.0
	for _, prediction := range predictions {
		sumProbabilities += prediction.WinProbability
	}

	if sumProbabilities != 100.0 {
		adjustmentFactor := 100.0 / sumProbabilities
		for _, prediction := range predictions {
			prediction.WinProbability *= adjustmentFactor
		}
	}

	return predictions
}
