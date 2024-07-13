package routes

import (
	controllers "football-league-simulation/controller"
	"net/http"

	_ "football-league-simulation/docs"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoutes() *httprouter.Router {
    router := httprouter.New()
    router.POST("/init", controllers.InitTeams)
    router.GET("/simulate", controllers.SimulateMatches)
		router.GET("/predict-champion-before-last-week", controllers.PredictChampionBeforeLastWeek)
    router.GET("/doc/*any", swaggerHandler)
    return router
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
     httpSwagger.WrapHandler(res, req)
}