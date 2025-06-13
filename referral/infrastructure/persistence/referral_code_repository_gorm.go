package persistence

import (
	"crypto/rand"
	"fmt"

	"internship/referral/domain/models"
	"internship/referral/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReferralCodeRepo struct {
	db *gorm.DB
}

func NewReferralCodeRepo(db *gorm.DB) repository.ReferralCodeRepository {
	return &ReferralCodeRepo{db}
}

func (r *ReferralCodeRepo) FindUserIDByCode(code string) (string, error) {
	var ref models.ReferralCodeModel
	if err := r.db.First(&ref, "code = ?", code).Error; err != nil {
		return "", fmt.Errorf("FindUserIDByCode failed: %w", err)
	}
	return ref.UserID.String(), nil
}

func (r *ReferralCodeRepo) GetReferralCodes(userID string) ([]models.ReferralCode, error) {
	var refs []models.ReferralCodeModel
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&refs).Error
	if err != nil {
		return nil, fmt.Errorf("GetReferralCodes failed: %w", err)
	}

	var out []models.ReferralCode
	for _, rc := range refs {
		out = append(out, models.ReferralCode{
			ID:     rc.ID,
			UserID: rc.UserID,
			Code:   rc.Code,
		})
	}
	return out, nil
}

func (r *ReferralCodeRepo) GenerateUniqueCode() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		code := make([]byte, 10)
		if _, err := rand.Read(code); err != nil {
			return "", fmt.Errorf("GenerateUniqueCode failed: %w", err)
		}
		for i := range code {
			code[i] = chars[int(code[i])%len(chars)]
		}
		var count int64
		r.db.Model(&models.ReferralCodeModel{}).Where("code = ?", string(code)).Count(&count)
		if count == 0 {
			return string(code), nil
		}
	}
}

func (r *ReferralCodeRepo) CreateReferralCode(userID string, code string) (models.ReferralCode, error) {
	id := uuid.New()
	ref := models.ReferralCodeModel{
		ID:     id,
		UserID: uuid.MustParse(userID),
		Code:   code,
	}
	if err := r.db.Create(&ref).Error; err != nil {
		return models.ReferralCode{}, fmt.Errorf("CreateReferralCode failed: %w", err)
	}
	return models.ReferralCode{
		ID:     ref.ID,
		UserID: ref.UserID,
		Code:   ref.Code,
	}, nil
}

func (r *ReferralCodeRepo) SyncReferralCode(userID string) error {
	refs, err := r.GetReferralCodes(userID)
	if err != nil {
		return fmt.Errorf("SyncReferralCode failed: %w", err)
	}
	if len(refs) == 0 {
		code, err := r.GenerateUniqueCode()
		if err != nil {
			return fmt.Errorf("SyncReferralCode failed to generate code: %w", err)
		}
		_, err = r.CreateReferralCode(userID, code)
		if err != nil {
			return fmt.Errorf("SyncReferralCode failed to create referral code: %w", err)
		}
	}
	return nil
}
