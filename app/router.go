package app

import (
	"ta-elearning/delivery"
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
	course := router.Group("/course")

	course.POST("/get", d.GetCourse)
	course.POST("/create", d.CreateCourse)
	course.POST("/update", d.UpdateCourse)
	course.POST("/delete", d.DeleteCourse)

	router.NoRoute(d.NoRoute)

	return router
}
