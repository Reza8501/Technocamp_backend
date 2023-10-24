package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) GetCourse(c context.Context, id []int, isFree string) ([]entity.Courses, error) {

	var data []entity.Courses

	db := r.mysqlConn.
		Model(&entity.Courses{})

	if len(id) > 0 {
		db = db.Where("id IN (?)", id)
	}

	if isFree != "" && isFree != "*" {
		db = db.Where("is_free = ?", isFree)
	}

	err := db.Scan(&data).Error

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, config.ErrRecordNotFound
	}

	return data, nil
}
