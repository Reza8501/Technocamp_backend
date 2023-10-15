package dto

type PayloadCourse struct {
	Id []string `json:"id" validate:"omitempty"`
}

type ReqCourseById struct {
	ID []int `json:"id"`
}

type ReqUpdateCourse struct {
	ID            int    `json:"id" validate:"required"`
	CourseTitle   string `json:"course_title" validate:"-"`
	CourseImage   string `json:"course_image" validate:"-"`
	CourseContent string `json:"course_content" validate:"-"`
	IsFree        bool   `json:"is_free" validate:"-"`
	IsActive      bool   `json:"is_active" validate:"-"`
	HasBuyer      bool   `json:"has_buyer" validate:"-"`
}
