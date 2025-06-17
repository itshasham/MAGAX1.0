package usecase

import (
	"errors"
	"internship/coupon/domain/models"
	"internship/coupon/domain/repository"
	"time"

	"github.com/google/uuid"
)

type CouponUsecase struct {
	repo repository.CouponRepository
}

func NewCouponUsecase(repo repository.CouponRepository) *CouponUsecase {
	return &CouponUsecase{repo: repo}
}

func (uc *CouponUsecase) Create(coupon *models.Coupon) error {
	return uc.repo.Create(coupon)
}

func (uc *CouponUsecase) GetAll(query map[string]interface{}, page, limit int) ([]models.Coupon, int64, error) {
	return uc.repo.FindAll(query, page, limit)
}

func (uc *CouponUsecase) GetByID(id uuid.UUID) (*models.Coupon, error) {
	return uc.repo.FindByID(id)
}

func (uc *CouponUsecase) Update(id uuid.UUID, updateData *models.Coupon) error {
	existing, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	updateData.ID = existing.ID
	return uc.repo.Update(updateData)
}

func (uc *CouponUsecase) Delete(id uuid.UUID) error {
	existing, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(existing)
}

func (uc *CouponUsecase) Validate(code string, orderAmount float64) (*models.Coupon, error) {
	coupon, err := uc.repo.FindByCode(code)
	if err != nil {
		return nil, errors.New("coupon not found")
	}

	now := time.Now()
	if !coupon.IsActive {
		return nil, errors.New("coupon is not active")
	}
	if coupon.StartDate != nil && now.Before(*coupon.StartDate) {
		return nil, errors.New("coupon not started yet")
	}
	if coupon.EndDate != nil && now.After(*coupon.EndDate) {
		return nil, errors.New("coupon expired")
	}
	if coupon.UsageLimit != nil && coupon.UsageCount >= *coupon.UsageLimit {
		return nil, errors.New("usage limit reached")
	}
	if coupon.MinOrderAmount != nil && orderAmount < *coupon.MinOrderAmount {
		return nil, errors.New("order amount too low")
	}

	return coupon, nil
}
