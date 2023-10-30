package repository

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/entity"
)

func (r *repository) ApproveTransaction(c context.Context, transactionCode string) error {

	trx := r.mysqlConn
	trx.Begin()

	errUpdate := trx.
		Model(&entity.UserTransaction{}).
		Where("transaction_code = ?", transactionCode).
		UpdateColumn("transaction_status", "success").
		Error
	if errUpdate != nil {
		trx.Rollback()
		return errUpdate
	}

	var data []entity.UserTransaction

	errGetItem := trx.Preload("User").Preload("Item").Preload("Item.Course").Where("transaction_code = ?", transactionCode).Find(&data).Error
	if errGetItem != nil {
		trx.Rollback()
		return errGetItem
	}

	if len(data) == 0 {
		trx.Rollback()
		return config.ErrRecordNotFound
	}

	for _, v := range data {
		for _, j := range v.Item {

			errInsertItem := trx.Create(&entity.UserHasCourse{
				UserId:   v.UserId,
				CourseId: j.CourseId,
			}).Error
			if errInsertItem != nil {
				trx.Rollback()
				return errInsertItem
			}
		}
	}

	trx.Commit()

	return nil
}
