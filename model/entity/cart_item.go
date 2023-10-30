package entity

type CartItem struct {
	Id       int           `json:"id" gorm:"column:id"`
	CartId   string        `json:"cartId" gorm:"column:cart_id"`
	CourseId int           `json:"courseId" gorm:"column:course_id"`
	Course   CoursesOutput `json:"course" gorm:"foreignkey:CourseId"`
}
