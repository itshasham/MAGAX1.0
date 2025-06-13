package repository

import model "internship/referral/domain/models"

type ReferralCodeRepository interface {
	FindUserIDByCode(code string) (string, error)
	GetReferralCodes(userID string) ([]model.ReferralCode, error)
	GenerateUniqueCode() (string, error)
	CreateReferralCode(userID string, code string) (model.ReferralCode, error)
	SyncReferralCode(userID string) error
}
