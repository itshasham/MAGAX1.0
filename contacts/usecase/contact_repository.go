package usecase

import "internship/contacts/domain"

type ContactRepository interface {
	Create(contact domain.Contact) (*domain.Contact, error)
}
