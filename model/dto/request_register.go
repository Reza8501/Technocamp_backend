package dto

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RequestRegisterUser struct {
	Id        string    `json:"id" validate:"required" gorm:"column:id"`
	FullName  string    `json:"fullName" validate:"required" gorm:"column:full_name"`
	Username  string    `json:"username" validate:"required" gorm:"column:username"`
	Email     string    `json:"email" validate:"required" gorm:"column:email"`
	Password  string    `json:"password" validate:"required" gorm:"column:password"`
	Role      string    `json:"role" validate:"omitempty" gorm:"column:role"`
	CreatedAt time.Time `json:"createdAt" validate:"omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty" gorm:"column:updated_at"`
}

type QueryParamRegisterVerification struct {
	Email string `form:"e" validate:"required"`
}

func (request *RequestRegisterUser) TransformRequest() {
	uuid, errUUID := uuid.NewUUID()
	if errUUID != nil {
		log.Print(errUUID.Error())
	}

	password := []byte(request.Password)

	// Generate a hashed password
	hashedPassword, errHashedPass := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if errHashedPass != nil {
		log.Print(errHashedPass)
	}

	request.Id = uuid.String()
	request.Role = "client"
	request.Password = string(hashedPassword)
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()
}
