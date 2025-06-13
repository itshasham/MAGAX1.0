package persistence

import (
	"errors"
	"fmt"
	models "internship/users/domain/models"
	"internship/users/domain/repository"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

// Implement the UserRepository interface methods here

func (r *userRepository) FindByWallet(wallet string) (*models.User, error) {
	var user models.User
	err := r.db.Where("wallet = ?", wallet).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // no user found is not an error
	}
	if err != nil {
		log.Printf("❌ FindByWallet failed: %v", err)
		return nil, fmt.Errorf("FindByWallet failed: %w", err)
	}
	return &user, nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Model(user).Updates(user).Error
}

func (r *userRepository) FindTopUsers(limit int) ([]models.User, error) {
	var users []models.User
	if err := r.db.Order("amount DESC").Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateBonus(id string, bonus float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).UpdateColumn("bonus", gorm.Expr("bonus + ?", bonus)).Error
}

func (r *userRepository) UpdateAmount(id string, amount float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).UpdateColumn("amount", gorm.Expr("amount + ?", amount)).Error
}
