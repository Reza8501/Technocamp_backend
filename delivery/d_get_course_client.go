package delivery

import (
	"net/http"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (d *delivery) GetCourseClient(c *gin.Context) {

	var req dto.ReqCourseById
	errBind := c.ShouldBindJSON(&req)

	if errBind != nil {
		errResp := response.BuildBadRequestResponse(response.ERROR_CODE_INVALID_DATA_TYPE, response.RESPONSE_CODE_BAD_REQUEST, response.RESPONSE_MESSAGE_INVALID_DATA_TYPE, errBind.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	userInfo := c.MustGet("userInfo").(jwt.MapClaims)
	if userInfo["role"].(string) != "client" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildUnauthorizedResponse("Invalid Authentication", "only client can access this api"))
		return
	}

	result := d.use.GetCourse(c, req)
	if *result.ErrorCode != response.ERROR_CODE_SUCCESS && (*result.ErrorCode == response.ERROR_CODE_DATA_NOT_FOUND || *result.ResponseMessage == response.RESPONSE_MESSAGE_DATA_NOT_FOUND) { // checking if error is data not found
		c.JSON(http.StatusNotFound, result)
		return
	} else if *result.ErrorCode == response.ERROR_CODE_GENERAL_ERROR { // general error from usecase
		c.JSON(http.StatusOK, result)
		return
	} else if *result.ErrorCode != response.ERROR_CODE_SUCCESS {
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	c.JSON(http.StatusOK, result)
}
