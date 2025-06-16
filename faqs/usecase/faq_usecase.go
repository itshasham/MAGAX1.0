package usecase

import (
	"internship/faqs/domain/models"
	domain "internship/faqs/domain/repository"
)

type FaqUsecase interface {
	GetPublicFaqs(category string, limit, page int) ([]models.Faq, error)
	GetFaqCategories() ([]string, error)
}

type faqUsecase struct {
	repo domain.FaqRepository
}

func NewFaqUsecase(repo domain.FaqRepository) FaqUsecase {
	return &faqUsecase{repo: repo}
}

func (f *faqUsecase) GetPublicFaqs(category string, limit, page int) ([]models.Faq, error) {
	return f.repo.GetPublicFaqs(category, limit, page)
}

func (f *faqUsecase) GetFaqCategories() ([]string, error) {
	return f.repo.GetFaqCategories()
}
