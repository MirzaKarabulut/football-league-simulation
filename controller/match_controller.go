package controllers

import (
	"fmt"
	"football-league-simulation/db"
	"football-league-simulation/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func InitTeams(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    teams := []models.Team{
        {Name: "Man City", Power: 90},
        {Name: "Real Madrid", Power: 85},
        {Name: "Barcelona", Power: 80},
        {Name: "Bayern Munich", Power: 88},
    }

    for _, team := range teams {
        _, err := db.DB.Exec("INSERT INTO teams (name, power, pts, w, d, l, gd) VALUES ($1, $2, 0, 0, 0, 0, 0)", team.Name, team.Power)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    w.WriteHeader(http.StatusCreated)
}

func SimulateMatches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    rand.Seed(time.Now().UnixNano())
    rows, err := db.DB.Query("SELECT id, name, power FROM teams")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var teams []models.Team
    for rows.Next() {
        var team models.Team
        err := rows.Scan(&team.ID, &team.Name, &team.Power)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        teams = append(teams, team)
    }

    if len(teams) < 2 {
        http.Error(w, "Not enough teams to simulate matches", http.StatusBadRequest)
        return
    }

    for week := 1; week <= 5; week++ {
        for i := 0; i < len(teams); i++ {
            for j := i + 1; j < len(teams); j++ {
                teamA := &teams[i]
                teamB := &teams[j]
                simulateMatch(teamA, teamB)
            }
        }
    }

    // Lig tablosunu görüntüleme
    displayLeagueTable()

    w.WriteHeader(http.StatusOK)
}

func simulateMatch(teamA, teamB *models.Team) {
    scoreA := rand.Intn(teamA.Power)
    scoreB := rand.Intn(teamB.Power)

    if scoreA > scoreB {
        teamA.Pts += 3
        teamA.W++
        teamB.L++
    } else if scoreA < scoreB {
        teamB.Pts += 3
        teamB.W++
        teamA.L++
    } else {
        teamA.Pts++
        teamB.Pts++
        teamA.D++
        teamB.D++
    }

    teamA.GD += (scoreA - scoreB)
    teamB.GD += (scoreB - scoreA)

    _, err := db.DB.Exec("UPDATE teams SET pts = $1, w = $2, d = $3, l = $4, gd = $5 WHERE id = $6", teamA.Pts, teamA.W, teamA.D, teamA.L, teamA.GD, teamA.ID)
    if err != nil {
        fmt.Println("Error updating teamA:", err)
    }

    _, err = db.DB.Exec("UPDATE teams SET pts = $1, w = $2, d = $3, l = $4, gd = $5 WHERE id = $6", teamB.Pts, teamB.W, teamB.D, teamB.L, teamB.GD, teamB.ID)
    if err != nil {
        fmt.Println("Error updating teamB:", err)
    }
}

func displayLeagueTable() {
    rows, err := db.DB.Query("SELECT name, pts, w, d, l, gd FROM teams ORDER BY pts DESC, gd DESC")
    if err != nil {
        fmt.Println("Error retrieving league table:", err)
        return
    }
    defer rows.Close()

    fmt.Println("League Table:")
    fmt.Println("Team\t\tPTS\tW\tD\tL\tGD")
    for rows.Next() {
        var name string
        var pts, w, d, l, gd int
        err := rows.Scan(&name, &pts, &w, &d, &l, &gd)
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return
        }
        fmt.Printf("%-15s %d\t%d\t%d\t%d\t%d\n", name, pts, w, d, l, gd)
    }
}
