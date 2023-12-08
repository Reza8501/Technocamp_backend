package entity

import (
	"fmt"
	"math/rand"
	"time"
)

type UserTransaction struct {
	Id                  string      `json:"id" gorm:"column:id"`
	UserId              string      `json:"userId" gorm:"column:user_id"`
	CartId              string      `json:"cartId" gorm:"column:cart_id"`
	TransactionCode     string      `json:"transactionCode" gorm:"column:transaction_code"`
	TransactionCurrency string      `json:"transactionCurrency" gorm:"column:transaction_currency"`
	TransactionTotal    float64     `json:"transactionTotal" gorm:"column:transaction_total"`
	TransactionMethod   string      `json:"transactionMethod" gorm:"column:transaction_method"`
	TransactionStatus   string      `json:"transactionStatus" gorm:"column:transaction_status"`
	TransactionProof    string      `json:"transactionProof" gorm:"column:transaction_proof"`
	CreatedAt           time.Time   `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt           time.Time   `json:"updatedAt" gorm:"column:updated_at"`
	User                UsersOutput `json:"user" gorm:"foreignkey:UserId"`
	Item                []CartItem  `json:"item" gorm:"foreignkey:CartId;association_foreignkey:CartId"`
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const codeLength = 4

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (u *UserTransaction) GenerateTransactionCode() {
	rand.Seed(time.Now().UnixNano())

	randomStr := randomString(codeLength)
	timestamp := time.Now().Unix()

	transactionCode := fmt.Sprintf("%s#%s", randomStr, fmt.Sprintf("%d", timestamp)[5:])
	u.TransactionCode = transactionCode
}
