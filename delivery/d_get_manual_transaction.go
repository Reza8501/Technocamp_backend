package delivery

import (
	"net/http"
	"ta-elearning/model/dto/response"

	"github.com/gin-gonic/gin"
)

func (d *delivery) GetManualTransaction(c *gin.Context) {

	result := d.use.GetManualTransaction(c)
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
