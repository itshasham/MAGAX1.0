package router

import (
	"internship/coupon/presentation/http"

	"github.com/gin-gonic/gin"
)

func RegisterCouponRoutes(r *gin.Engine, h *http.CouponHandler) {
	// Admin routes
	admin := r.Group("/admin/coupons")
	{
		admin.POST("", h.Create)       // POST /admin/coupons
		admin.GET("", h.FindAll)       // GET /admin/coupons
		admin.GET("/:id", h.FindOne)   // GET /admin/coupons/:id
		admin.PATCH("/:id", h.Update)  // PUT /admin/coupons/:id
		admin.DELETE("/:id", h.Delete) // DELETE /admin/coupons/:id
	}

	// Public route for coupon validation
	r.POST("/coupons/validate", h.Validate) // POST /coupons/validate
}
