package router

import (
	httpHandler "internship/faqs/presentation/http"

	"github.com/gin-gonic/gin"
)

func RegisterFaqRoutes(r *gin.Engine, handler *httpHandler.FaqHandler) {
	r.GET("/faqs", handler.GetPublicFaqs)
	r.GET("/faqs/categories", handler.GetFaqCategories)
}
