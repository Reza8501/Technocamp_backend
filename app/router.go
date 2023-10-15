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

	router := gin.New()

	router.NoRoute(d.NoRoute)
	router.GET("/ping", d.Ping)

	return router
}
