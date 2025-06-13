package usecase

import (
	"fmt"
	refRepo "internship/referral/domain/repository"
	userModel "internship/users/domain/models"
	userRepo "internship/users/domain/repository"
	"internship/users/infrastructure/token"
	"log"

	"github.com/google/uuid"
)

type UserUseCase struct {
	UserRepo userRepo.UserRepository
	RefRepo  refRepo.ReferralCodeRepository
}

func (uc *UserUseCase) FindByID(id string) (*userModel.User, error) {
	user, err := uc.UserRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by ID [%s]: %w", id, err)
	}
	return user, nil
}

func (uc *UserUseCase) ConnectWallet(wallet string, ref *string) (*userModel.User, string, error) {
	user, err := uc.UserRepo.FindByWallet(wallet)
	if err != nil {
		log.Printf("❌ FindByWallet failed: %v", err)
		return nil, "", fmt.Errorf("find wallet error: %w", err)
	}

	if user == nil {
		log.Println("ℹ️ Wallet not found. Creating new user.")
		newUser := &userModel.User{
			ID:         uuid.New(),
			Wallet:     wallet,
			ReferralBy: ref,
		}
		if err := uc.UserRepo.Create(newUser); err != nil {
			log.Printf("❌ Failed to create user: %v", err)
			return nil, "", fmt.Errorf("user create error: %w", err)
		}
		user = newUser
	}

	log.Println("ℹ️ Syncing referral code.")
	if err := uc.RefRepo.SyncReferralCode(user.ID.String()); err != nil {
		log.Printf("❌ SyncReferralCode failed: %v", err)
		return nil, "", fmt.Errorf("referral sync error: %w", err)
	}

	log.Println("ℹ️ Generating JWT token.")
	tok, err := token.Generate(user.ID.String())
	if err != nil {
		log.Printf("❌ Token generation failed: %v", err)
		return nil, "", fmt.Errorf("token error: %w", err)
	}

	return user, tok, nil
}

func (uc *UserUseCase) UpdateUser(u *userModel.User) error {
	if err := uc.UserRepo.Update(u); err != nil {
		return fmt.Errorf("failed to update user [%s]: %w", u.ID, err)
	}
	return nil
}

func (uc *UserUseCase) GetTopUsers() ([]userModel.User, error) {
	users, err := uc.UserRepo.FindTopUsers(10)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top users: %w", err)
	}
	return users, nil
}

func (uc *UserUseCase) AddBonus(id string, bonus float64) error {
	if err := uc.UserRepo.UpdateBonus(id, bonus); err != nil {
		return fmt.Errorf("failed to update bonus for user [%s]: %w", id, err)
	}
	return nil
}

func (uc *UserUseCase) AddAmount(id string, amount float64) error {
	if err := uc.UserRepo.UpdateAmount(id, amount); err != nil {
		return fmt.Errorf("failed to update amount for user [%s]: %w", id, err)
	}
	return nil
}
