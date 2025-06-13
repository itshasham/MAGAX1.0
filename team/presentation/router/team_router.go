package router

import (
	"internship/team/presentation/http"

	"github.com/gin-gonic/gin"
)

func RegisterTeamRoutes(r *gin.Engine, handler *http.TeamHandler) {
	team := r.Group("/teams")
	{
		team.GET("/", handler.GetTeams)         // GET    /teams?limit=10&page=1
		team.GET("/:id", handler.GetTeamByID)   // GET    /teams/:id
		team.POST("/", handler.CreateTeam)      // POST   /teams
		team.PUT("/:id", handler.UpdateTeam)    // PUT    /teams/:id
		team.DELETE("/:id", handler.DeleteTeam) // DELETE /teams/:id
	}
}
