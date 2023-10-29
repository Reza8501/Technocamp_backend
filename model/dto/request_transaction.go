package dto

type RequestTransaction struct {
	CartId string `json:"cartId" validate:"required"`
}

type TotalCoursePrice struct {
	Total int `json:"total" gorm:"total"`
}

type ApproveTransaction struct {
	TransactionCode string `json:"transactionCode" validate:"required"`
}
