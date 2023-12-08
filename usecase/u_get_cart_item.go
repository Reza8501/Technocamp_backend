package usecase

import (
	"context"
	"ta-elearning/config"
	"ta-elearning/model/dto/response"
	"ta-elearning/model/entity"

	"github.com/jinzhu/gorm"
)

func (usecase *usecase) GetCartItem(c context.Context, userId string) *response.ResponseContainer {
	userCart, errUserCart := usecase.Repo.GetUserCartByUserId(c, userId)
	if errUserCart != nil && gorm.IsRecordNotFoundError(errUserCart) {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error getting cart", errUserCart.Error())
	}

	cartItem, errCartItem := usecase.Repo.GetCartItem(c, userCart.Id)
	if errCartItem != nil && !gorm.IsRecordNotFoundError(errCartItem) {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error getting cart item", errCartItem.Error())
	}

	var coursesId []int
	if len(cartItem) > 0 {
		for _, v := range cartItem {
			coursesId = append(coursesId, v.CourseId)
		}

		courses, errCourses := usecase.Repo.GetCourse(c, coursesId, "", "", true)
		if errCourses != nil && errCourses.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errCourses != nil {
			return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, response.RESPONSE_MESSAGE_DATABASE_ERROR, errCourses.Error())
		} else if errCourses != nil && gorm.IsRecordNotFoundError(errCourses) {
			return response.BuildDataNotFoundResponse()
		}

		return response.BuildSuccessResponse(&entity.ResponseCartItem{
			CartId: userCart.Id,
			Course: courses,
		})
	}

	return response.BuildSuccessResponse(&entity.ResponseCartItem{
		CartId: userCart.Id,
	})
}
