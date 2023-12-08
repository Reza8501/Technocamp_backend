package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) DeleteCartItem(c context.Context, cartId string, courseId int) error {
	var data entity.CartItem
	errDb := repository.mysqlConn.
		Model(&data).
		Where("cart_id = ?", cartId).
		Where("course_id = ?", courseId).
		Delete(&data).Error
	if errDb != nil {
		return errDb
	}

	return nil
}
