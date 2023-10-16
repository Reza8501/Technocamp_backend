package repository

import (
	"context"
	"ta-elearning/model/dto"
)

func (repository *repository) CreateCourse(c context.Context, req dto.ReqCreateCourse) error {

	err := repository.mysqlConn.
		Create(&req).Error
	if err != nil {
		return err
	}

	return nil
}
