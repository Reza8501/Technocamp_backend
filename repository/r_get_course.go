package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) GetCourse(c context.Context, id []int, isFree string, IsActive string, HasBuyer bool) ([]entity.Courses, error) {

	var data []entity.Courses

	db := r.mysqlConn.
		Model(&entity.Courses{}).Preload("Buyer")

	if len(id) > 0 {
		db = db.Where("id IN (?)", id)
	}

	if isFree != "" && isFree != "*" {
		db = db.Where("is_free = ?", isFree)
	}

	if IsActive != "" && IsActive != "*" {
		if IsActive == "1" {
			db = db.Where("is_active = ?", true)
		} else {
			db = db.Where("is_active = ?", false)
		}
	}

	err := db.Find(&data).Error

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, config.ErrRecordNotFound
	}

	return data, nil
}
