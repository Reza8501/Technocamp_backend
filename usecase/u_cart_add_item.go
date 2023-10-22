package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"
	"ta-elearning/model/entity"
)

func (u *usecase) CartAddItem(c context.Context, req dto.RequestCartItem, userId string) *response.ResponseContainer {
	var userCart *entity.UserCart
	var errUserCart, errCreateUser error
	userCart, errUserCart = u.Repo.GetUserCartByUserId(c, userId)
	if errUserCart != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error getting cart", errUserCart.Error())
	}

	if userCart == nil {
		userCart, errCreateUser = u.Repo.CreateUserCart(c, userId)
		if errCreateUser != nil {
			return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error creating cart", errCreateUser.Error())
		}
	}

	errCreateItem := u.Repo.CreateCartItem(c, userCart.Id, req.CourseId)
	if errCreateItem != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error creating cart item", errCreateItem.Error())
	}

	return response.BuildSuccessResponse(nil)
}
