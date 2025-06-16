package usecase

import "internship/contacts/domain"

type ContactUsecase interface {
	CreateContact(contact domain.Contact) (*domain.Contact, error)
}

type contactUsecase struct {
	repo ContactRepository
}

func NewContactUsecase(r ContactRepository) ContactUsecase {
	return &contactUsecase{repo: r}
}

func (u *contactUsecase) CreateContact(contact domain.Contact) (*domain.Contact, error) {
	return u.repo.Create(contact)
}
