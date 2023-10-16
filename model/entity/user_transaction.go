package entity

import "time"

type UserTransaction struct {
	Id                  string    `json:"id" gorm:"column:id"`
	UserId              string    `json:"userId" gorm:"column:user_id"`
	CartId              string    `json:"cartId" gorm:"column:cart_id"`
	TransactionCode     string    `json:"transactionCode" gorm:"column:transaction_code"`
	TransactionCurrency string    `json:"transactionCurrency" gorm:"column:transaction_currency"`
	TransactionTotal    float64   `json:"transactionTotal" gorm:"column:transaction_total"`
	TransactionMethod   string    `json:"transactionMethod" gorm:"column:transaction_method"`
	TransactionStatus   string    `json:"transactionStatus" gorm:"column:transaction_status"`
	CreatedAt           time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt           time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
