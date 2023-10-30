package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) GetManualTransaction(c context.Context) ([]entity.UserTransaction, error) {

	var data []entity.UserTransaction
	db := r.mysqlConn.Preload("User").Preload("Item").Preload("Item.Course").Find(&data)
	if db.Error != nil {
		return nil, db.Error
	}

	if len(data) == 0 {
		return nil, config.ErrRecordNotFound
	}

	return data, nil
}
