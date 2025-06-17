package persistence

import (
	"internship/blog/domain/models"
	"internship/blog/domain/repository"

	"gorm.io/gorm"
)

type blogPGRepo struct {
	db *gorm.DB
}

func NewBlogPGRepository(db *gorm.DB) repository.BlogRepository {
	return &blogPGRepo{db}
}

func (r *blogPGRepo) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *blogPGRepo) FindAll(search, category string, limit, offset int) ([]models.Blog, int64, error) {
	var blogs []models.Blog
	var count int64
	query := r.db.Model(&models.Blog{}).Where("status = ?", "enabled")
	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}
	err := query.Count(&count).Limit(limit).Offset(offset).Order("sort_order ASC, created_at DESC").Find(&blogs).Error
	return blogs, count, err
}
func (r *blogPGRepo) Delete(blog *models.Blog) error {
	return r.db.Delete(blog).Error
}
func (r *blogPGRepo) FindRecent(limit int) ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Where("status = ?", "enabled").Order("created_at DESC").Limit(limit).Find(&blogs).Error
	return blogs, err
}

func (r *blogPGRepo) FindBySlug(slug string) (*models.Blog, error) {
	var blog models.Blog
	err := r.db.Where("slug = ? AND status = ?", slug, "enabled").First(&blog).Error
	return &blog, err
}

func (r *blogPGRepo) FindByID(id int) (*models.Blog, error) {
	var blog models.Blog
	err := r.db.First(&blog, id).Error
	return &blog, err
}

func (r *blogPGRepo) Update(blog *models.Blog) error {
	return r.db.Save(blog).Error
}
