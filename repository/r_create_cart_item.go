package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) CreateCartItem(c context.Context, cartId string, courseId int) error {

	// Database insert
	err := repository.mysqlConn.
		Model(&entity.CartItem{}).
		Create(&entity.CartItem{
			CartId:   cartId,
			CourseId: courseId,
		}).Error

	return err
}
