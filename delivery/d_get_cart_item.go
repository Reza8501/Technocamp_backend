package delivery

import (
	"net/http"
	"ta-elearning/model/dto/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (d *delivery) GetCartItem(c *gin.Context) {
	userInfo := c.MustGet("userInfo").(jwt.MapClaims)
	result := d.use.GetCartItem(c.Request.Context(), userInfo["id"].(string))
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
