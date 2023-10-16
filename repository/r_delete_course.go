package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) DeleteCourse(c context.Context, id []int) error {

	var data entity.Courses

	if len(id) == 0 {
		return nil
	}

	deleteResult := repository.mysqlConn.
		Model(&data).
		Where("id IN (?)", id).
		Delete(&data)

	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
