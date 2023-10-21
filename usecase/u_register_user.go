package usecase

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"ta-elearning/model/dto"
	"ta-elearning/model/dto/response"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func (u *usecase) RegisterUser(c context.Context, req dto.RequestRegisterUser) *response.ResponseContainer {

	user, errUser := u.Repo.GetUserByUsername(c, req.Username)
	if errUser != nil && !gorm.IsRecordNotFoundError(errUser) {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error getting user", errUser.Error())
	}

	userByEmail, errUserByEmail := u.Repo.GetUserByUsername(c, req.Email)
	if errUserByEmail != nil && !gorm.IsRecordNotFoundError(errUserByEmail) {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error getting user", errUserByEmail.Error())
	}

	if user != nil && userByEmail != nil {
		return response.BuildBadRequestResponse(response.ERROR_CODE_INVALID_BODY_REQUEST, response.RESPONSE_CODE_BAD_REQUEST, "Data already exist", "User has been registered")
	}

	errRegister := u.Repo.RegisterUser(c, req)
	if errRegister != nil && gorm.IsRecordNotFoundError(errRegister) {
		return response.BuildDataNotFoundResponse()
	} else if errRegister != nil {
		return response.BuildInternalErrorResponse(response.ERROR_CODE_DATABASE_ERROR, response.RESPONSE_CODE_INTERNAL_ERROR, "Error registering user", errRegister.Error())
	}

	go func() {
		err := SendVerificationEmail(req.Email)
		if err != nil {
			log.Print(err.Error())
		} else {
			log.Print("Successfully send verification email")
		}
	}()

	return response.BuildSuccessResponse(nil)
}

func SendVerificationEmail(email string) error {
	auth := smtp.PlainAuth("", viper.GetString("SMTP_USERNAME"), viper.GetString("SMTP_PASSWORD"), viper.GetString("SMTP_HOST"))

	// build message
	message := fmt.Sprintf("From: %s\r\n", viper.GetString("APP_EMAIL"))
	message += fmt.Sprintf("To: %s\r\n", email)
	message += fmt.Sprintf("Subject: %s\r\n", "Testing")
	message += fmt.Sprintf("\r\n%s\r\n", "Testing")

	err := smtp.SendMail(fmt.Sprintf("%s:%s", viper.GetString("SMTP_HOST"), viper.GetString("SMTP_PORT")), auth, viper.GetString("APP_EMAIL"), []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
