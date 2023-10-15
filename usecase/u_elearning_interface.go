package usecase

import "ta-elearning/repository"

type Usecase interface {
	Ping() string
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{
		repo: r,
	}
}

func (u *usecase) Ping() string {
	return u.repo.Ping()
}
