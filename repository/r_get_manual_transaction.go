package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) GetManualTransaction(c context.Context) ([]entity.UserTransaction, error) {

	var data []entity.UserTransaction

	db := r.mysqlConn.
		Model(&entity.UserTransaction{}).Where("transaction_method = ?", "manual")

	err := db.Scan(&data).Error

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, config.ErrRecordNotFound
	}

	return data, nil
}
