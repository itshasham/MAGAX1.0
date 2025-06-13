package repository

import model "internship/users/domain/models"

type UserRepository interface {
	FindByWallet(wallet string) (*model.User, error)
	FindByID(id string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	UpdateBonus(id string, bonus float64) error
	UpdateAmount(id string, amount float64) error
	FindTopUsers(limit int) ([]model.User, error)
}
