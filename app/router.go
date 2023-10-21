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

	router.NoRoute(d.NoRoute)

	return router
}
