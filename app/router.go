package app

import (
	"ta-elearning/delivery"
	"ta-elearning/middleware"
	"ta-elearning/repository"
	"ta-elearning/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	r := repository.NewRepository(mysqlConn)
	u := usecase.NewUsecase(r)
	d := delivery.NewDelivery(u)

	router := gin.Default()

	// authentication segment
	auth := router.Group("/auth")
	auth.POST("/login", d.Login)

	// course segment
	course := router.Group("/course")
	course.Use(middleware.MiddlewareAuth())
	{
		course.POST("/get", d.GetCourse)
		course.POST("/create", d.CreateCourse)
		course.POST("/update", d.UpdateCourse)
		course.POST("/delete", d.DeleteCourse)
	}

	// users segment
	users := router.Group("/users")
	users.POST("/register", d.RegisterUser)
	users.GET("/register/verification", d.RegisterVerification)

	// cart segment
	cart := router.Group("/cart")
	cart.Use(middleware.MiddlewareAuth())
	{
		cart.POST("/add-item", d.CartAddItem)
	}

	// transaction segment
	transaction := router.Group("/transaction")
	transaction.Use(middleware.MiddlewareAuth())
	{
		transaction.POST("/", d.Transaction)
		transaction.POST("/manual", d.GetManualTransaction)
		transaction.POST("/client", d.GetClientTransaction)
		transaction.POST("/approve", d.ApproveTransaction)
	}

	router.NoRoute(d.NoRoute)

	return router
}
