package router

import (
	"internship/contacts/presentation/http"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterContactRoutes(r *gin.Engine, handler *http.ContactHandler) {
	routeGroup := r.Group("/")
	{
		log.Println("🔗 Registering POST /contacts route")
		routeGroup.POST("/contacts", func(c *gin.Context) {
			log.Println("➡️  Incoming request: POST /contacts")
			handler.CreateContact(c)
		})
	}

	for _, route := range r.Routes() {
		if route.Path == "/contacts" {
			log.Printf("✅ Route registered: [%s] %s", route.Method, route.Path)
		}
	}
}
