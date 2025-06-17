package router

import (
	"internship/contacts/presentation/http"

	"github.com/gin-gonic/gin"
)

func RegisterContactRoutes(r *gin.Engine, handler *http.ContactHandler) {
	routeGroup := r.Group("/")
	{

		routeGroup.POST("/contacts", func(c *gin.Context) {

			handler.CreateContact(c)
		})
	}

	for _, route := range r.Routes() {
		if route.Path == "/contacts" {

		}
	}
}
