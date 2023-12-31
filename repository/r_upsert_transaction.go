package repository

import (
	"context"
	"ta-elearning/model/entity"
	"time"

	"github.com/google/uuid"
)

func (repository *repository) UpsertTransaction(c context.Context, userId string, cartId string, total int) error {

	var request entity.UserTransaction

	uuid, errUUID := uuid.NewUUID()
	if errUUID != nil {
		return errUUID
	}

	request.Id = uuid.String()
	request.CartId = cartId
	request.UserId = userId
	request.TransactionCurrency = "IDR"
	request.TransactionMethod = "manual"
	request.TransactionStatus = "pending"
	request.TransactionTotal = float64(total)
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
	request.GenerateTransactionCode()

	trx := repository.mysqlConn
	// trx.Begin()

	err := trx.Create(&request).Error
	if err != nil {
		// trx.Rollback()
		return err
	}

	errDelete := trx.Delete(&entity.UserCart{}, "user_id = ?", userId).Error
	if errDelete != nil {
		// trx.Rollback()
		return errDelete
	}
	// trx.Commit()

	return nil
}
