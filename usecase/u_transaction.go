package usecase

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) Transaction(c context.Context, req dto.RequestTransaction, userId string) *response.ResponseContainer {

	resultTotalPrice, errTotalPrice := u.Repo.GetTotalPrice(c, req.CartId)
	if errTotalPrice != nil && errTotalPrice.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errTotalPrice != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, response.RESPONSE_MESSAGE_DATABASE_ERROR, errTotalPrice.Error())
	} else if errTotalPrice != nil && gorm.IsRecordNotFoundError(errTotalPrice) {
		return response.BuildDataNotFoundResponse()
	}

	err := u.Repo.UpsertTransaction(c, userId, req.CartId, resultTotalPrice.Total)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, response.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	} else if err != nil && gorm.IsRecordNotFoundError(err) {
		return response.BuildDataNotFoundResponse()
	}

	return response.BuildSuccessResponse(nil)
}
