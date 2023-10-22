package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) CreateUserCart(c context.Context, userId string) (*entity.UserCart, error) {

	request := &entity.UserCart{
		UserId: userId,
	}
	request.TransformInsert()
	var data entity.UserCart
	// Database insert
	err := repository.mysqlConn.
		Model(&entity.UserCart{}).
		Create(request).Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, err
}
