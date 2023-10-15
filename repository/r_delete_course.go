package repository

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/entity"
)

func (repository *repository) DeleteCourse(c context.Context, req dto.ReqCourseById) error {

	var data entity.Courses

	deleteResult := repository.mysqlConn.
		Model(&data).
		Where("id = ?", req.ID).
		Delete(&data)

	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
