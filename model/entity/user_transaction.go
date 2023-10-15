package entity

import "time"

type UserTransaction struct {
	Id                  string    `json:"id" gorm:"column:id"`
	UserId              string    `json:"user_id" gorm:"column:user_id"`
	CartId              string    `json:"cart_id" gorm:"column:cart_id"`
	TransactionCode     string    `json:"transaction_code" gorm:"column:transaction_code"`
	TransactionCurrency string    `json:"transaction_currency" gorm:"column:transaction_currency"`
	TransactionTotal    float64   `json:"transaction_total" gorm:"column:transaction_total"`
	TransactionMethod   string    `json:"transaction_method" gorm:"column:transaction_method"`
	TransactionStatus   string    `json:"transaction_status" gorm:"column:transaction_status"`
	CreatedAt           time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"column:updated_at"`
}
