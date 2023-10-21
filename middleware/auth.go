package middleware

import (
	"errors"
	"net/http"
	"strings"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildUnauthorizedResponse("Invalid Authentication", "Invalid Auth Token"))
			return
		}

		tokenString := strings.Replace(authorization, "Bearer ", "", -1)
		token, errToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("signing method invalid")
			} else if method != dto.JWT_SIGNING_METHOD {
				return nil, errors.New("signing method invalid")
			}

			return []byte(viper.GetString("APP_NAME")), nil
		})
		if errToken != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildUnauthorizedResponse("Invalid Authentication", errToken.Error()))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.BuildUnauthorizedResponse("Invalid Authentication", "Invalid Auth Token"))
			return
		}
		c.Set("userInfo", claims)
		c.Next()
	}
}
