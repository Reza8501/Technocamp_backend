package entity

type Users struct {
	Id       string `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
}
