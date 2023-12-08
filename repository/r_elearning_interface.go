package repository

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/entity"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Ping() string
	GetCourse(c context.Context, id []int, isFree string, IsActive string, HasBuyer bool) ([]entity.Courses, error)
	CreateCourse(c context.Context, req dto.ReqCreateCourse) error
	UpdateCourse(c context.Context, req dto.ReqUpdateCourse) error
	DeleteCourse(c context.Context, id []int) error
	GetUserByUsername(c context.Context, username string) (*entity.Users, error)
	RegisterUser(c context.Context, req dto.RequestRegisterUser) error
	UpdateStatusVerification(c context.Context, email string, status bool) error
	CreateUserCart(c context.Context, userId string) (*entity.UserCart, error)
	GetCartItem(c context.Context, cartId string) ([]entity.CartItem, error)
	DeleteCartItem(c context.Context, cartId string, courseId int) error
	GetUserCartByUserId(c context.Context, userId string) (*entity.UserCart, error)
	CreateCartItem(c context.Context, cartId string, courseId int) error
	GetTotalPrice(c context.Context, cartId string) (*dto.TotalCoursePrice, error)
	UpsertTransaction(c context.Context, userId string, cartId string, total int) error
	GetManualTransaction(c context.Context) ([]entity.UserTransaction, error)
	GetClientTransaction(c context.Context, userId string) ([]entity.UserTransaction, error)
	ApproveTransaction(c context.Context, transactionCode string) error
	RejectTransaction(c context.Context, transactionCode string) error
	ProofTransaction(c context.Context, transactionCode string, proof string) error
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
