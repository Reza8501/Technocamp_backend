package dto

type ReqCourseById struct {
	ID []int `json:"id" validate:"omitempty"`
}

type ReqDeleteCourseById struct {
	ID []int `json:"id" validate:"required"`
}

type ReqUpdateCourse struct {
	Id            int    `json:"id" gorm:"column:id"`
	CourseTitle   string `json:"course_title" gorm:"column:course_title"`
	CourseImage   string `json:"course_image" gorm:"column:course_image"`
	CourseContent string `json:"course_content" gorm:"column:course_content"`
	IsFree        bool   `json:"is_free" gorm:"column:is_free"`
	IsActive      bool   `json:"is_active" gorm:"column:is_active"`
	HasBuyer      bool   `json:"has_buyer" gorm:"column:has_buyer"`
}

type ReqCreateCourse struct {
	Id            int    `json:"id" gorm:"column:id"`
	CourseTitle   string `json:"course_title" gorm:"column:course_title"`
	CourseImage   string `json:"course_image" gorm:"column:course_image"`
	CourseContent string `json:"course_content" gorm:"column:course_content"`
	IsFree        bool   `json:"is_free" gorm:"column:is_free"`
	IsActive      bool   `json:"is_active" gorm:"column:is_active"`
	HasBuyer      bool   `json:"has_buyer" gorm:"column:has_buyer"`
}