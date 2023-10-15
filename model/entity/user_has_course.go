package entity

type UserHasCourse struct {
	UserId   string `json:"user_id" gorm:"column:user_id"`
	CourseId int    `json:"course_id" gorm:"column:course_id"`
}
