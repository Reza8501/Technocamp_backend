package entity

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type UserCart struct {
	Id        string     `json:"id" gorm:"column:id"`
	UserId    string     `json:"userId" gorm:"column:user_id"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

func (request *UserCart) TransformInsert() {
	uuid, errUUID := uuid.NewUUID()
	if errUUID != nil {
		log.Print(errUUID.Error())
	}

	request.Id = uuid.String()
}
