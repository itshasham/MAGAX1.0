package usecase

import (
	"internship/blog/domain/models"
	"internship/blog/domain/repository"
)

type BlogUsecase struct {
	repo repository.BlogRepository
}

func NewBlogUsecase(r repository.BlogRepository) *BlogUsecase {
	return &BlogUsecase{repo: r}
}

func (uc *BlogUsecase) Create(blog *models.Blog) error {
	return uc.repo.Create(blog)
}

func (uc *BlogUsecase) Update(blog *models.Blog) error {
	return uc.repo.Update(blog)
}

func (uc *BlogUsecase) Delete(id int) error {
	blog, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(blog)
}

func (uc *BlogUsecase) FindAll(search, category string, page, limit int) ([]models.Blog, int64, error) {
	offset := (page - 1) * limit
	return uc.repo.FindAll(search, category, limit, offset)
}

func (uc *BlogUsecase) FindRecent(limit int) ([]models.Blog, error) {
	return uc.repo.FindRecent(limit)
}

func (uc *BlogUsecase) FindBySlug(slug string) (*models.Blog, error) {
	blog, err := uc.repo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	blog.Views++
	_ = uc.repo.Update(blog)
	return blog, nil
}

func (uc *BlogUsecase) FindByID(id int) (*models.Blog, error) {
	return uc.repo.FindByID(id)
}
