package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) CreateCourse(c context.Context, req entity.Courses) error {

	err := repository.mysqlConn.
		Create(&req).Error
	if err != nil {
		return err
	}

	return nil
}
