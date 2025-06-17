package repository

import (
	"internship/coupon/domain/models"

	"github.com/google/uuid"
)

type CouponRepository interface {
	Create(coupon *models.Coupon) error
	FindAll(query map[string]interface{}, page, limit int) ([]models.Coupon, int64, error)
	FindByID(id uuid.UUID) (*models.Coupon, error)
	Update(coupon *models.Coupon) error
	Delete(coupon *models.Coupon) error
	FindByCode(code string) (*models.Coupon, error)
}
