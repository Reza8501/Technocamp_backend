package entity

const (
	TABLE_USERS            = "users"
	TABLE_USER_TRANSACTION = "user_transaction"
	TABLE_USER_HAS_COURSE  = "user_has_course"
	TABLE_USER_CART        = "user_cart"
	TABLE_ROLES            = "roles"
	TABLE_COURSES          = "courses"
	TABLE_CART_ITEM        = "cart_item"
)

type Tabler interface {
	TableName() string
}

func (Users) TableName() string {
	return TABLE_USERS
}
func (UsersOutput) TableName() string {
	return TABLE_USERS
}
func (UserTransaction) TableName() string {
	return TABLE_USER_TRANSACTION
}
func (UserHasCourse) TableName() string {
	return TABLE_USER_HAS_COURSE
}
func (UserCart) TableName() string {
	return TABLE_USER_CART
}
func (Roles) TableName() string {
	return TABLE_ROLES
}
func (Courses) TableName() string {
	return TABLE_COURSES
}
func (CoursesOutput) TableName() string {
	return TABLE_COURSES
}
func (CartItem) TableName() string {
	return TABLE_CART_ITEM
}
