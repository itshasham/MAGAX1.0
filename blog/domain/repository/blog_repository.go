package repository

import (
	"internship/blog/domain/models"
)

type BlogRepository interface {
	Create(blog *models.Blog) error
	FindAll(search, category string, limit, offset int) ([]models.Blog, int64, error)
	FindRecent(limit int) ([]models.Blog, error)
	FindBySlug(slug string) (*models.Blog, error)
	FindByID(id int) (*models.Blog, error)
	Update(blog *models.Blog) error
	Delete(blog *models.Blog) error
}
