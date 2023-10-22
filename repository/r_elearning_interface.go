package repository

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/entity"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Ping() string
	GetCourse(c context.Context, id []int) ([]entity.Courses, error)
	CreateCourse(c context.Context, req dto.ReqCreateCourse) error
	UpdateCourse(c context.Context, req dto.ReqUpdateCourse) error
	DeleteCourse(c context.Context, id []int) error
	GetUserByUsername(c context.Context, username string) (*entity.Users, error)
	RegisterUser(c context.Context, req dto.RequestRegisterUser) error
	UpdateStatusVerification(c context.Context, email string, status bool) error
	CreateUserCart(c context.Context, userId string) (*entity.UserCart, error)
	GetUserCartByUserId(c context.Context, userId string) (*entity.UserCart, error)
	CreateCartItem(c context.Context, cartId string, courseId int) error
}

type repository struct {
	mysqlConn *gorm.DB
}

func NewRepository(mysqlConn *gorm.DB) Repository {
	return &repository{
		mysqlConn: mysqlConn,
	}
}

func (r *repository) Ping() string {
	return "pong"
}
