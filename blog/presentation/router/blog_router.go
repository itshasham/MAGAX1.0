package router

import (
	handler "internship/blog/presentation/http"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(r *gin.Engine, h *handler.BlogHandler) {
	blogRoutes := r.Group("/blogs")
	{
		blogRoutes.GET("", h.FindAll)           // GET /blogs
		blogRoutes.GET("/recent", h.FindRecent) // GET /blogs/recent
		blogRoutes.GET("/:slug", h.FindOne)     // GET /blogs/:slug

		blogRoutes.POST("", h.Create)       // POST /blogs
		blogRoutes.PATCH("/:id", h.Update)  // PUT /blogs/:id
		blogRoutes.DELETE("/:id", h.Delete) // DELETE /blogs/:id
	}
}
