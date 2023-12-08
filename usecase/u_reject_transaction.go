package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) RejectTransaction(c context.Context, request dto.RejectTransaction) *response.ResponseContainer {
	errRejectTransaction := u.Repo.RejectTransaction(c, request.TransactionCode)
	if errRejectTransaction != nil && gorm.IsRecordNotFoundError(errRejectTransaction) {
		return response.BuildDataNotFoundResponse()
	} else if errRejectTransaction != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error reject transaction", errRejectTransaction.Error())
	}

	return response.BuildSuccessResponse(nil)
}
