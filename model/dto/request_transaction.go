package dto

type RequestTransaction struct {
	CartId string `json:"cartId" validate:"required"`
}

type TotalCoursePrice struct {
	Total int `json:"total" gorm:"total"`
}
