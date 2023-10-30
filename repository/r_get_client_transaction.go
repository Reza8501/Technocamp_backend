package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) GetClientTransaction(c context.Context, userId string) ([]entity.UserTransaction, error) {

	var data []entity.UserTransaction

	db := r.mysqlConn.Preload("User").Preload("Item").Preload("Item.Course").Where("user_id = ?", userId).Find(&data)
	if db.Error != nil {
		return nil, db.Error
	}

	if len(data) == 0 {
		return nil, config.ErrRecordNotFound
	}

	return data, nil
}
