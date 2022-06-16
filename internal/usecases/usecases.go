package usecases

import (
	"testTask/internal/domain"
	"testTask/internal/repository"
)

type UseCase struct {
	repo *repository.Repository
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) GetByIDs(ids []int) ([]*domain.Prop, error) {
	result, err := u.repo.GetByIDs(ids)
	return result, err
}

func (u *UseCase) CreateProps(props domain.Props) error {
	return u.repo.InsertProps(props)
}

func (u *UseCase) UpdateProps(props domain.Props) error {
	return u.repo.UpdateProps(props)
}

func (u *UseCase) CreateProp(props domain.Props) error {
	return u.repo.InsertProp(props)
}
