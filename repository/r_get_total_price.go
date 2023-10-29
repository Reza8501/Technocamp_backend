package repository

import (
	"context"
	"fmt"
	"ta-elearning/model/dto"
	"ta-elearning/model/entity"
)

func (r *repository) GetTotalPrice(c context.Context, cartId string) (*dto.TotalCoursePrice, error) {

	var data dto.TotalCoursePrice

	db := r.mysqlConn.
		Model(&entity.CartItem{}).Select("SUM(B.price) AS total").Joins(fmt.Sprintf("JOIN ta.courses AS B ON cart_item.course_id = B.id WHERE cart_item.cart_id = '%s'", cartId))

	err := db.Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}
