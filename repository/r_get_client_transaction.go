package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) GetClientTransaction(c context.Context, userId string) ([]entity.UserTransaction, error) {

	var data []entity.UserTransaction

	db := r.mysqlConn.
		Model(&entity.UserTransaction{}).Where("user_id = ?", userId)

	err := db.Scan(&data).Error

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, config.ErrRecordNotFound
	}

	return data, nil
}
