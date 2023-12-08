package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) GetUserCartByUserId(c context.Context, userId string) (*entity.UserCart, error) {

	var data entity.UserCart
	err := repository.mysqlConn.
		Model(&entity.UserCart{}).
		Where("user_id = ?", userId).Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
