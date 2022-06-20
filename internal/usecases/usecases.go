package usecases

import (
	"testTask/internal/domain"
)

type repositorer interface{
	GetByIDs([]uint)([]*domain.Prop, error)
	InsertProps(*domain.Props)error
	UpdateProps(*domain.Props)error
}

type UseCase struct {
	repo repositorer
}

func NewUseCase(repo repositorer) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) GetByIDs(ids []uint) ([]*domain.Prop, error) {
	result, err := u.repo.GetByIDs(ids)
	return result, err
}

func (u *UseCase) CreateProps(props *domain.Props) error {
	return u.repo.InsertProps(props)
}

func (u *UseCase) UpdateProps(props *domain.Props) error {
	return u.repo.UpdateProps(props)
}

// func (u *UseCase) CreateProp(props *domain.Props) error {
// 	return u.repo.InsertProp(props)
// }
