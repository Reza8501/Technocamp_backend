package dto

type RequestCartItem struct {
	CourseId int `json:"courseId" gorm:"column:course_id"`
}
