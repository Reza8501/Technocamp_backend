package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (r *repository) GetUserByUsername(c context.Context, username string) (*entity.Users, error) {
	var data entity.Users
	errDb := r.mysqlConn.
		Model(&entity.Users{}).
		Where("username = ?", username).
		Or("email = ?", username).
		Limit(1).
		Scan(&data).
		Error

	if errDb != nil {
		return nil, errDb
	}

	return &data, nil
}
