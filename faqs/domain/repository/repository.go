package repository

import "internship/faqs/domain/models"

type FaqRepository interface {
	GetPublicFaqs(category string, limit, page int) ([]models.Faq, error)
	GetFaqCategories() ([]string, error)
}
