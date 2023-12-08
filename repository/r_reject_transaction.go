package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (r *repository) RejectTransaction(c context.Context, transactionCode string) error {

	trx := r.mysqlConn
	trx.Begin()

	errUpdate := trx.
		Model(&entity.UserTransaction{}).
		Where("transaction_code = ?", transactionCode).
		UpdateColumn("transaction_status", "rejected").
		Error
	if errUpdate != nil {
		trx.Rollback()
		return errUpdate
	}
	trx.Commit()

	return nil
}
