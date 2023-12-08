package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (usecase *usecase) ProofTransaction(c context.Context, request dto.RequestProofTransaction, userId string) *response.ResponseContainer {
	errProof := usecase.Repo.ProofTransaction(c, request.TransactionCode, request.Proof)
	if errProof != nil && !gorm.IsRecordNotFoundError(errProof) {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error sent proof transaction", errProof.Error())
	}

	return response.BuildSuccessResponse(nil)
}
