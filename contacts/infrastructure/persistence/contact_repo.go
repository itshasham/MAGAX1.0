package persistence

import (
	"internship/contacts/domain"
	"internship/contacts/usecase"

	"gorm.io/gorm"
)

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) usecase.ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) Create(contact domain.Contact) (*domain.Contact, error) {
	err := r.db.Create(&contact).Error
	return &contact, err
}
