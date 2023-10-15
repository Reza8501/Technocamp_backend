package entity

type CartItem struct {
	Id       int    `json:"id" gorm:"column:id"`
	CartId   string `json:"cart_id" gorm:"column:cart_id"`
	CourseId int    `json:"course_id" gorm:"column:course_id"`
}
