package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Ping() string
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
