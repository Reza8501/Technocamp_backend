package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (r *repository) ApproveTransaction(c context.Context, transactionCode string) error {
	errDb := r.mysqlConn.
		Model(&entity.UserTransaction{}).
		Where("transaction_code = ?", transactionCode).
		UpdateColumn("transaction_status", "success").
		Error

	if errDb != nil {
		return errDb
	}

	return nil
}
