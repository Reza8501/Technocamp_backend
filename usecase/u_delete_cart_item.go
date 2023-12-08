package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (usecase *usecase) DeleteCartItem(c context.Context, request dto.RequestDeleteCartItem, userId string) *response.ResponseContainer {
	errDelete := usecase.Repo.DeleteCartItem(c, request.CartId, request.CourseId)
	if errDelete != nil && !gorm.IsRecordNotFoundError(errDelete) {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error deleting cart item", errDelete.Error())
	}

	return response.BuildSuccessResponse(nil)
}
