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
	Login(c context.Context, request dto.RequestLogin) *response.ResponseContainer
	RegisterUser(c context.Context, req dto.RequestRegisterUser) *response.ResponseContainer
	RegisterVerification(c context.Context, email string) *response.ResponseContainer
	CartAddItem(c context.Context, req dto.RequestCartItem, userId string) *response.ResponseContainer
	GetCartItem(c context.Context, userId string) *response.ResponseContainer
	DeleteCartItem(c context.Context, request dto.RequestDeleteCartItem, userId string) *response.ResponseContainer
	Transaction(c context.Context, req dto.RequestTransaction, userId string) *response.ResponseContainer
	GetManualTransaction(c context.Context) *response.ResponseContainer
	GetClientTransaction(c context.Context, userId string) *response.ResponseContainer
	ApproveTransaction(c context.Context, request dto.ApproveTransaction) *response.ResponseContainer
	RejectTransaction(c context.Context, request dto.RejectTransaction) *response.ResponseContainer
	ProofTransaction(c context.Context, request dto.RequestProofTransaction, userId string) *response.ResponseContainer
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
