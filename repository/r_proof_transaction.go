package repository

import (
	"context"
	"ta-elearning/model/entity"
)

func (repository *repository) ProofTransaction(c context.Context, transactionCode string, proof string) error {
	errDb := repository.mysqlConn.
		Model(&entity.UserTransaction{}).
		Where("transaction_code = ?", transactionCode).
		Updates(&entity.UserTransaction{
			TransactionProof: proof,
		}).Error
	if errDb != nil {
		return errDb
	}

	return nil
}
