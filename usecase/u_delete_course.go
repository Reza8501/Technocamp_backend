package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) DeleteCourse(c context.Context, req dto.ReqDeleteCourseById) *response.ResponseContainer {

	if len(req.ID) == 0 {
		return response.BuildDataNotFoundResponse()
	}

	course, errGetById := u.Repo.GetCourse(c, req.ID)
	if errGetById != nil && gorm.IsRecordNotFoundError(errGetById) {
		return response.BuildDataNotFoundResponseWithMessage("updated failed, id might not be found")
	} else if errGetById != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error updating Course", errGetById.Error())
	}

	if len(course) > 0 {
		for _, v := range course {
			if v.HasBuyer {
				return response.BuildBadRequestResponse(response.ERROR_CODE_INVALID_BODY_REQUEST, response.RESPONSE_CODE_BAD_REQUEST, "Error deleting Course", "This Course Has Buyer")
			}
		}
	}

	errDelete := u.Repo.DeleteCourse(c, req.ID)
	if errDelete != nil && gorm.IsRecordNotFoundError(errDelete) {
		return response.BuildDataNotFoundResponseWithMessage("delete failed, id might not be found")
	} else if errDelete != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error deleting Course", errDelete.Error())
	}

	return response.BuildSuccessResponse(nil)
}
