package entity

import "time"

type UserCart struct {
	Id        string    `json:"id" gorm:"column:id"`
	UserId    string    `json:"user_id" gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
