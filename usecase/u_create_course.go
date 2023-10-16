package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
)

func (u *usecase) CreateCourse(c context.Context, req dto.ReqCreateCourse) *response.ResponseContainer {

	errCreate := u.Repo.CreateCourse(c, req)
	if errCreate != nil && gorm.IsRecordNotFoundError(errCreate) {
		return response.BuildDataNotFoundResponse()
	} else if errCreate != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error creating Course", errCreate.Error())
	}

	return response.BuildSuccessResponse(nil)
}
