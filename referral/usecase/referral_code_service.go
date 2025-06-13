package usecase

import (
	model "internship/referral/domain/models"
	"internship/referral/domain/repository"
)

type ReferralCodeService struct {
	Repo repository.ReferralCodeRepository
}

func (s *ReferralCodeService) FindUserByCode(code string) (string, error) {
	return s.Repo.FindUserIDByCode(code)
}

func (s *ReferralCodeService) GetReferralCodes(userID string) ([]model.ReferralCode, error) {
	return s.Repo.GetReferralCodes(userID)
}

func (s *ReferralCodeService) CreateNewReferralCode(userID string) (model.ReferralCode, error) {
	code, err := s.Repo.GenerateUniqueCode()
	if err != nil {
		return model.ReferralCode{}, err
	}
	return s.Repo.CreateReferralCode(userID, code)
}

func (s *ReferralCodeService) SyncReferralCode(userID string) error {
	return s.Repo.SyncReferralCode(userID)
}
