package http

import (
	"fmt"
	"internship/faqs/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FaqHandler struct {
	Usecase usecase.FaqUsecase
}

func NewFaqHandler(uc usecase.FaqUsecase) *FaqHandler {
	return &FaqHandler{Usecase: uc}
}

func (h *FaqHandler) GetPublicFaqs(c *gin.Context) {
	category := c.DefaultQuery("category", "")
	limit := parseIntWithDefault(c.DefaultQuery("limit", "10"), 10)
	page := parseIntWithDefault(c.DefaultQuery("page", "1"), 1)

	faqs, err := h.Usecase.GetPublicFaqs(category, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FAQs"})
		return
	}
	c.JSON(http.StatusOK, faqs)
}

func (h *FaqHandler) GetFaqCategories(c *gin.Context) {
	categories, err := h.Usecase.GetFaqCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func parseIntWithDefault(value string, def int) int {
	var i int
	_, err := fmt.Sscanf(value, "%d", &i)
	if err != nil {
		return def
	}
	return i
}
