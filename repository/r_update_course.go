package repository

import (
	"context"
	"ta-elearning/model/dto"
)

func (repository *repository) UpdateCourse(c context.Context, req dto.ReqUpdateCourse) error {

	result := repository.mysqlConn.
		Model(&req).
		Where("id = ?", req.Id).
		Updates(&req)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
