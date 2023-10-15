package entity

type Roles struct {
	Name string `json:"name" gorm:"column:name"`
}
