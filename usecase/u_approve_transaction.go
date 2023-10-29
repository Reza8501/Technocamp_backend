package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) ApproveTransaction(c context.Context, request dto.ApproveTransaction) *response.ResponseContainer {
	errApproveTransaction := u.Repo.ApproveTransaction(c, request.TransactionCode)
	if errApproveTransaction != nil && gorm.IsRecordNotFoundError(errApproveTransaction) {
		return response.BuildDataNotFoundResponse()
	} else if errApproveTransaction != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error approve transaction", errApproveTransaction.Error())
	}

	return response.BuildSuccessResponse(nil)
}
