package repository

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/entity"
)

func (r *repository) RegisterUser(c context.Context, req dto.RequestRegisterUser) error {
	req.TransformRequest()

	errDb := r.mysqlConn.Model(&entity.Users{}).Create(&req).
		Error

	if errDb != nil {
		return errDb
	}

	return nil
}
