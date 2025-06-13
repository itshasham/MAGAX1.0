package usecase

import "internship/team/domain"

type TeamRepository interface {
	FindAndCountEnabled(limit, offset int) ([]domain.Team, int, error)
	FindAndCountBySearch(search string, limit, offset int) ([]domain.Team, int, error)
	FindByID(id int) (*domain.Team, error)
	Create(team domain.Team) (*domain.Team, error)
	Update(id int, team domain.Team) (*domain.Team, error)
	Delete(id int) error
}
