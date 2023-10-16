package entity

type UserHasCourse struct {
	UserId   string `json:"userId" gorm:"column:user_id"`
	CourseId int    `json:"courseId" gorm:"column:course_id"`
}
