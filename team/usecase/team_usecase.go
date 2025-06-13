package usecase

import (
	"errors"
	"internship/team/domain"
)

type TeamUsecase interface {
	GetEnabledTeams(limit, page int) ([]domain.Team, int, error)
	GetAdminTeams(search string, limit, page int) ([]domain.Team, int, error)
	GetTeamByID(id int) (*domain.Team, error)
	CreateTeam(team domain.Team) (*domain.Team, error)
	UpdateTeam(id int, team domain.Team) (*domain.Team, error)
	DeleteTeam(id int) error
}

type teamUsecase struct {
	repo TeamRepository
}

func NewTeamUsecase(repo TeamRepository) TeamUsecase {
	return &teamUsecase{repo: repo}
}

func (u *teamUsecase) GetEnabledTeams(limit, page int) ([]domain.Team, int, error) {
	offset := (page - 1) * limit
	return u.repo.FindAndCountEnabled(limit, offset)
}

func (u *teamUsecase) GetAdminTeams(search string, limit, page int) ([]domain.Team, int, error) {
	offset := (page - 1) * limit
	return u.repo.FindAndCountBySearch(search, limit, offset)
}

func (u *teamUsecase) GetTeamByID(id int) (*domain.Team, error) {
	team, err := u.repo.FindByID(id)
	if err != nil || team == nil {
		return nil, errors.New("team not found")
	}
	return team, nil
}

func (u *teamUsecase) CreateTeam(team domain.Team) (*domain.Team, error) {
	return u.repo.Create(team)
}

func (u *teamUsecase) UpdateTeam(id int, team domain.Team) (*domain.Team, error) {
	return u.repo.Update(id, team)
}

func (u *teamUsecase) DeleteTeam(id int) error {
	return u.repo.Delete(id)
}
