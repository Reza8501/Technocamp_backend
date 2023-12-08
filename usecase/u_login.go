package usecase

import (
	"context"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(c context.Context, request dto.RequestLogin) *response.ResponseContainer {
	// check user exists
	user, errUser := u.Repo.GetUserByUsername(c, request.Username)
	if errUser != nil && gorm.IsRecordNotFoundError(errUser) {
		return response.BuildUnauthorizedResponse("Invalid Authentication", "Invalid Username or Password")
	} else if errUser != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error getting user", errUser.Error())
	}

	// comparing password
	validPassword := ComparePassword(request.Password, user.Password)
	if !validPassword {
		return response.BuildUnauthorizedResponse("Invalid Authentication", "Invalid Username or Password")
	}

	// check status verification
	if !user.StatusVerification {
		return response.BuildUnauthorizedResponse("Invalid Authentication", "User not verified")
	}

	// jwt claims
	jwtClaims := dto.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: viper.GetString("APP_NAME"),
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(dto.JWT_EXPIRATION_DURATION),
			},
		},
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}

	// jwt get token
	token := jwt.NewWithClaims(dto.JWT_SIGNING_METHOD, jwtClaims)
	signedToken, errSignedToken := token.SignedString([]byte(viper.GetString("APP_NAME")))
	if errSignedToken != nil {
		return response.BuildUnauthorizedResponse("Invalid Authentication", errSignedToken.Error())
	}

	return response.BuildSuccessResponse(&dto.ResponseJwt{
		Token: signedToken,
		Role:  user.Role,
		Name:  user.FullName,
	})
}

func ComparePassword(requestPassword, sourcePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(sourcePassword), []byte(requestPassword))
	if err == nil { // Passwords match
		return true
	} else if err == bcrypt.ErrMismatchedHashAndPassword { // Passwords do not match
		return false
	} else { // Handle other errors
		return false
	}
}
