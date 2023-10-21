package usecase

import (
	"context"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) RegisterVerification(c context.Context, email string) *response.ResponseContainer {
	errVerification := u.Repo.UpdateStatusVerification(c, email, true)
	if errVerification != nil && gorm.IsRecordNotFoundError(errVerification) {
		return response.BuildDataNotFoundResponse()
	} else if errVerification != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error verification user", errVerification.Error())
	}

	return response.BuildSuccessResponse(nil)
}
