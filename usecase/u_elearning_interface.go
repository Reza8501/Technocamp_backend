package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"
	"ta-elearning/repository"
)

type Usecase interface {
	Ping() string
	GetCourse(c context.Context, req dto.ReqCourseById) *response.ResponseContainer
	CreateCourse(c context.Context, reqData dto.ReqCreateCourse) *response.ResponseContainer
	UpdateCourse(c context.Context, req dto.ReqUpdateCourse) *response.ResponseContainer
	DeleteCourse(c context.Context, req dto.ReqDeleteCourseById) *response.ResponseContainer
}

type usecase struct {
	Repo repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{
		Repo: r,
	}
}

func (u *usecase) Ping() string {
	return u.Repo.Ping()
}
