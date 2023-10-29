package usecase

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) GetClientTransaction(c context.Context, userId string) *response.ResponseContainer {

	resultCourse, err := u.Repo.GetClientTransaction(c, userId)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, response.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	} else if err != nil && gorm.IsRecordNotFoundError(err) {
		return response.BuildDataNotFoundResponse()
	}

	return response.BuildSuccessResponse(resultCourse)
}
