package router

import (
	handler "internship/users/presentation/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/wallet-connect", userHandler.ConnectWallet)
		userGroup.GET("/me/:id", userHandler.GetUserByID)
		userGroup.PUT("/me/:id", userHandler.UpdateUser)
		userGroup.GET("/leaderboard", userHandler.Leaderboard)
	}
}
