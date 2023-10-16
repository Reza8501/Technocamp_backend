package delivery

import (
	"errors"
	"io"
	"net/http"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/gin-gonic/gin"
)

func (d *delivery) CreateCourse(c *gin.Context) {

	var req dto.ReqCreateCourse
	errBind := c.ShouldBindJSON(&req)

	if errBind != nil && errors.Is(errBind, io.EOF) { // checking if body req is empty
		errResp := response.BuildBadRequestResponse(response.ERROR_CODE_BODY_REQUEST_EMPTY, response.RESPONSE_CODE_BAD_REQUEST, response.RESPONSE_MESSAGE_BODY_REQ_EMPTY, errBind.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	} else if errBind != nil {
		errResp := response.BuildBadRequestResponse(response.ERROR_CODE_INVALID_DATA_TYPE, response.RESPONSE_CODE_BAD_REQUEST, response.RESPONSE_MESSAGE_INVALID_DATA_TYPE, errBind.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	result := d.use.CreateCourse(c.Request.Context(), req)
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
	return
}
