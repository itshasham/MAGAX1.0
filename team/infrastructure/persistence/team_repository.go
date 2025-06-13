package persistence

import (
	"internship/team/domain"
	"internship/team/usecase"

	"gorm.io/gorm"
)

type teamRepo struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) usecase.TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) FindAndCountEnabled(limit, offset int) ([]domain.Team, int, error) {
	var teams []domain.Team
	var count int64
	err := r.db.Model(&domain.Team{}).Where("status = ?", domain.StatusEnabled).
		Order("sort_order ASC").Count(&count).Limit(limit).Offset(offset).Find(&teams).Error
	return teams, int(count), err
}

func (r *teamRepo) FindAndCountBySearch(search string, limit, offset int) ([]domain.Team, int, error) {
	var teams []domain.Team
	var count int64
	query := r.db.Model(&domain.Team{}).Where("status = ?", domain.StatusEnabled)
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	err := query.Order("sort_order ASC").Count(&count).Limit(limit).Offset(offset).Find(&teams).Error
	return teams, int(count), err
}

func (r *teamRepo) FindByID(id int) (*domain.Team, error) {
	var team domain.Team
	err := r.db.First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *teamRepo) Create(team domain.Team) (*domain.Team, error) {
	err := r.db.Create(&team).Error
	return &team, err
}

func (r *teamRepo) Update(id int, team domain.Team) (*domain.Team, error) {
	err := r.db.Model(&domain.Team{}).Where("id = ?", id).Updates(team).Error
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *teamRepo) Delete(id int) error {
	return r.db.Delete(&domain.Team{}, id).Error
}
