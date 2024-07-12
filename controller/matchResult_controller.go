package controllers

import (
	"fmt"
	"football-league-simulation/db"
	"football-league-simulation/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SaveMatchResult(week int, teamA, teamB string, scoreA, scoreB int) error {
    _, err := db.DB.Exec("INSERT INTO match_results (week, team_a, team_b, score_a, score_b) VALUES ($1, $2, $3, $4, $5)", week, teamA, teamB, scoreA, scoreB)
    return err
}

func GetMatchResultsByWeek(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    week := ps.ByName("week")

    rows, err := db.DB.Query("SELECT team_a, team_b, score_a, score_b FROM match_results WHERE week = $1", week)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    fmt.Printf("Match results for week %s:\n", week)
    for rows.Next() {
        var result models.MatchResult
        err := rows.Scan(&result.TeamA, &result.TeamB, &result.ScoreA, &result.ScoreB)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Printf("Match: %s %d - %d %s\n", result.TeamA, result.ScoreA, result.ScoreB, result.TeamB)
    }

    w.WriteHeader(http.StatusOK)
}
