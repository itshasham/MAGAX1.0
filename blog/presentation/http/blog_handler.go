package http

import (
	"fmt"
	"internship/blog/domain/models"
	"internship/blog/usecase"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	uc *usecase.BlogUsecase
}

func NewBlogHandler(uc *usecase.BlogUsecase) *BlogHandler {
	return &BlogHandler{uc: uc}
}

func (h *BlogHandler) FindAll(c *gin.Context) {
	search := c.Query("search")
	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	blogs, count, err := h.uc.FindAll(search, category, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blogs, "count": count})
}

func (h *BlogHandler) Create(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&blog); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusConflict, gin.H{"error": "Slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog"})
		return
	}

	c.JSON(http.StatusCreated, blog)
}

// PUT /blogs/:id
func (h *BlogHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Step 1: Load the existing blog
	existing, err := h.uc.FindByID(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// Step 2: Parse input
	var input models.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 3: Only update fields that are set
	if input.Title != "" {
		existing.Title = input.Title
	}
	if input.Slug != "" {
		existing.Slug = input.Slug
	}
	if input.Content != "" {
		existing.Content = input.Content
	}
	if input.FeaturedImg != "" {
		existing.FeaturedImg = input.FeaturedImg
	}
	if input.Status != "" {
		existing.Status = input.Status
	}
	if input.Category != "" {
		existing.Category = input.Category
	}
	if input.Author != "" {
		existing.Author = input.Author
	}
	if input.MetaTitle != "" {
		existing.MetaTitle = input.MetaTitle
	}
	if input.MetaDescription != "" {
		existing.MetaDescription = input.MetaDescription
	}
	if input.MetaKeywords != "" {
		existing.MetaKeywords = input.MetaKeywords
	}
	if input.ReadTime != 0 {
		existing.ReadTime = input.ReadTime
	}
	if input.SortOrder != 0 {
		existing.SortOrder = input.SortOrder
	}

	existing.UpdatedAt = time.Now()

	// Step 4: Save to DB
	if err := h.uc.Update(existing); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

// DELETE /blogs/:id
func (h *BlogHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	existing, err := h.uc.FindByID(id)
	if err != nil || existing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	// ✅ Developer-friendly confirmation
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Blog with ID %d deleted successfully", id),
	})
}

func (h *BlogHandler) FindRecent(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	blogs, err := h.uc.FindRecent(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) FindOne(c *gin.Context) {
	slug := c.Param("slug")
	blog, err := h.uc.FindBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}
