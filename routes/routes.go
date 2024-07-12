package routes

import (
	controllers "football-league-simulation/controller"

	"github.com/julienschmidt/httprouter"
)

func InitRoutes() *httprouter.Router {
    router := httprouter.New()
    router.POST("/init", controllers.InitTeams)
    router.POST("/simulate", controllers.SimulateMatches)
		router.GET("/predict-champion-before-last-week", controllers.PredictChampionBeforeLastWeek)
		router.GET("/get-match-results/:week", controllers.GetMatchResultsByWeek)
    return router
}
