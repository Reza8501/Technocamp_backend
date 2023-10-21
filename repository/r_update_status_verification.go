package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (r *repository) UpdateStatusVerification(c context.Context, email string, status bool) error {
	errDb := r.mysqlConn.
		Model(&entity.Users{}).
		Where("email = ?", email).
		Update(&entity.Users{
			StatusVerification: status,
		}).
		Error

	if errDb != nil {
		return errDb
	}

	return nil
}
