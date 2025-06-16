package persistence

import (
	"internship/faqs/domain/models"
	domain "internship/faqs/domain/repository"

	"gorm.io/gorm"
)

type faqRepository struct {
	db *gorm.DB
}

func NewFaqRepository(db *gorm.DB) domain.FaqRepository {
	return &faqRepository{db: db}
}

func (r *faqRepository) GetPublicFaqs(category string, limit, page int) ([]models.Faq, error) {
	var faqs []models.Faq
	offset := (page - 1) * limit
	query := r.db.Where("status = ?", "enabled")
	if category != "" {
		query = query.Where("category = ?", category)
	}
	err := query.Order("sort_order asc").Limit(limit).Offset(offset).Find(&faqs).Error
	return faqs, err
}

func (r *faqRepository) GetFaqCategories() ([]string, error) {
	var categories []string
	err := r.db.Model(&models.Faq{}).Distinct().Pluck("category", &categories).Error
	return categories, err
}
