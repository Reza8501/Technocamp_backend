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
