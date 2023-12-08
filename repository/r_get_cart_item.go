package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) GetCartItem(c context.Context, cartId string) ([]entity.CartItem, error) {
	var data []entity.CartItem
	errDb := repository.mysqlConn.
		Model(&entity.CartItem{}).
		Where("cart_id = ?", cartId).
		Scan(&data).Error

	if errDb != nil {
		return data, errDb
	}

	return data, nil
}
