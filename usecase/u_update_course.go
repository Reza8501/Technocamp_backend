package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) UpdateCourse(c context.Context, req dto.ReqUpdateCourse) *response.ResponseContainer {

	_, errGetById := u.Repo.GetCourse(c, []int{req.ID})
	if errGetById != nil && gorm.IsRecordNotFoundError(errGetById) {
		return response.BuildDataNotFoundResponseWithMessage("updated failed, id might not be found")
	} else if errGetById != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error updating Course", errGetById.Error())
	}

	errUpdate := u.Repo.UpdateCourse(c, req)
	if errUpdate != nil && gorm.IsRecordNotFoundError(errUpdate) {
		return response.BuildDataNotFoundResponseWithMessage("updated failed, id might not be found")
	} else if errUpdate != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error updating Course", errUpdate.Error())
	}

	return response.BuildSuccessResponse(errUpdate)
}
