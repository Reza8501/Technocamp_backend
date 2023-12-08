package dto

type RequestCartItem struct {
	CourseId int `json:"courseId" gorm:"column:course_id"`
}

type RequestDeleteCartItem struct {
	CartId   string `json:"cartId" gorm:"column:cart_id"`
	CourseId int    `json:"courseId" gorm:"column:course_id"`
}
