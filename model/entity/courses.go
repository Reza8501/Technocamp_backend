package entity

type Courses struct {
	Id            int    `json:"id" gorm:"column:id"`
	CourseTitle   string `json:"courseTitle" gorm:"column:course_title"`
	CourseImage   string `json:"courseImage" gorm:"column:course_image"`
	CourseContent string `json:"courseContent" gorm:"column:course_content"`
	IsFree        bool   `json:"isFree" gorm:"column:is_free"`
	IsActive      bool   `json:"isActive" gorm:"column:is_active"`
	HasBuyer      bool   `json:"hasBuyer" gorm:"column:has_buyer"`
}

type CoursesOutput struct {
	Id            int    `json:"id" gorm:"column:id"`
	CourseTitle   string `json:"courseTitle" gorm:"column:course_title"`
	CourseImage   string `json:"courseImage" gorm:"column:course_image"`
	CourseContent string `json:"courseContent" gorm:"column:course_content"`
}
