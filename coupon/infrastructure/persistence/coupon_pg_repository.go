package persistence

import (
	"internship/coupon/domain/models"
	"internship/coupon/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CouponPGRepo struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) repository.CouponRepository {
	return &CouponPGRepo{db: db}
}

func (r *CouponPGRepo) Create(coupon *models.Coupon) error {
	return r.db.Create(coupon).Error
}

func (r *CouponPGRepo) FindAll(filters map[string]interface{}, page, limit int) ([]models.Coupon, int64, error) {
	var coupons []models.Coupon
	var count int64

	offset := (page - 1) * limit
	query := r.db.Model(&models.Coupon{}).Where(filters)

	query.Count(&count)
	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&coupons).Error

	return coupons, count, err
}

func (r *CouponPGRepo) FindByID(id uuid.UUID) (*models.Coupon, error) {
	var coupon models.Coupon
	err := r.db.First(&coupon, "id = ?", id).Error
	return &coupon, err
}

func (r *CouponPGRepo) Update(coupon *models.Coupon) error {
	return r.db.Save(coupon).Error
}

func (r *CouponPGRepo) Delete(coupon *models.Coupon) error {
	return r.db.Delete(coupon).Error
}

func (r *CouponPGRepo) FindByCode(code string) (*models.Coupon, error) {
	var coupon models.Coupon
	err := r.db.Where("code = ?", code).First(&coupon).Error
	return &coupon, err
}
