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
		dto.SMTP_AUTH = smtp.PlainAuth("", viper.GetString("SMTP_USERNAME"), viper.GetString("SMTP_PASSWORD"), viper.GetString("SMTP_HOST"))
		templateData := struct {
			Name string
			URL  string
		}{
			Name: req.FullName,
			URL:  fmt.Sprintf("http://localhost:8080/users/register/verification?e=%s", req.Email),
		}

		r := dto.NewRequestEmail(
			[]string{req.Email},
			"Verification Account E-Learning",
			"",
		)

		if err := r.ParseTemplate(fmt.Sprintf("%s/verification.html", dto.PATH_TEMPLATE_VERIFICATION_MAIL), templateData); err == nil {
			if _, errSendMail := r.SendEmail(); errSendMail != nil {
				log.Print(errSendMail.Error())
			} else {
				log.Print("Successfully send verification email")
			}
		} else {
			log.Print("Error parsing email template")
		}
	}()

	return response.BuildSuccessResponse(nil)
}
