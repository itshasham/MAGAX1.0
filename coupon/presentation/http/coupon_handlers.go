package http

import (
	"internship/coupon/domain/models"
	"internship/coupon/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	uc *usecase.CouponUsecase
}

func NewCouponHandler(uc *usecase.CouponUsecase) *CouponHandler {
	return &CouponHandler{uc: uc}
}

func (h *CouponHandler) Create(c *gin.Context) {
	var input models.Coupon
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coupon"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func (h *CouponHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	filters := make(map[string]interface{})
	if c.Query("active") == "true" {
		filters["is_active"] = true
	}
	if c.Query("expired") == "true" {
		filters["end_date <"] = "NOW()" // custom handling if needed
	}

	coupons, total, err := h.uc.GetAll(filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch coupons"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": coupons, "total": total, "page": page, "limit": limit})
}

func (h *CouponHandler) FindOne(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	coupon, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}
	c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	// Load the existing coupon
	existing, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}

	var input models.Coupon
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Merge logic: only update fields provided
	if input.Code != "" {
		existing.Code = input.Code
	}
	existing.IsActive = input.IsActive
	if input.StartDate != nil {
		existing.StartDate = input.StartDate
	}
	if input.EndDate != nil {
		existing.EndDate = input.EndDate
	}
	if input.UsageLimit != nil {
		existing.UsageLimit = input.UsageLimit
	}
	if input.MinOrderAmount != nil {
		existing.MinOrderAmount = input.MinOrderAmount
	}
	if input.UsageCount != 0 {
		existing.UsageCount = input.UsageCount
	}

	existing.UpdatedAt = time.Now()

	// Save
	if err := h.uc.Update(id, existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update coupon"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func (h *CouponHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coupon deleted successfully"})
}
func (h *CouponHandler) Validate(c *gin.Context) {
	var body struct {
		Code        string  `json:"code"`
		OrderAmount float64 `json:"orderAmount"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coupon, err := h.uc.Validate(body.Code, body.OrderAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupon)
}
