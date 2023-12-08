package dto

type ReqCourseById struct {
	ID       []int  `json:"id" validate:"omitempty"`
	IsFree   string `json:"isFree" validate:"omitempty"`
	IsActive string `json:"isActive" validate:"omitempty"`
	HasBuyer bool   `json:"hasBuyer" validate:"omitempty"`
}

type ReqDeleteCourseById struct {
	ID []int `json:"id" validate:"required"`
}

type ReqUpdateCourse struct {
	Id            int    `json:"id" gorm:"column:id"`
	CourseTitle   string `json:"courseTitle" gorm:"column:course_title"`
	CourseImage   string `json:"courseImage" gorm:"column:course_image"`
	CourseContent string `json:"courseContent" gorm:"column:course_content"`
	IsFree        bool   `json:"isFree" gorm:"column:is_free"`
	IsActive      bool   `json:"isActive" gorm:"column:is_active"`
	HasBuyer      bool   `json:"hasBuyer" gorm:"column:has_buyer"`
	Price         int    `json:"price" gorm:"column:price"`
}

type ReqCreateCourse struct {
	Id            int    `json:"id" gorm:"column:id"`
	CourseTitle   string `json:"courseTitle" gorm:"column:course_title"`
	CourseImage   string `json:"courseImage" gorm:"column:course_image"`
	CourseContent string `json:"courseContent" gorm:"column:course_content"`
	IsFree        bool   `json:"isFree" gorm:"column:is_free"`
	IsActive      bool   `json:"isActive" gorm:"column:is_active"`
	HasBuyer      bool   `json:"hasBuyer" gorm:"column:has_buyer"`
	Price         int    `json:"price" gorm:"column:price"`
}
