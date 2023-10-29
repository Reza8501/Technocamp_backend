package delivery

import (
	"ta-elearning/usecase"

	"github.com/gin-gonic/gin"
)

type Delivery interface {
	Ping(c *gin.Context)
	NoRoute(c *gin.Context)
	GetCourse(c *gin.Context)
	CreateCourse(c *gin.Context)
	UpdateCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
	Login(c *gin.Context)
	RegisterUser(c *gin.Context)
	RegisterVerification(c *gin.Context)
	CartAddItem(c *gin.Context)
	Transaction(c *gin.Context)
	GetManualTransaction(c *gin.Context)
	GetClientTransaction(c *gin.Context)
	ApproveTransaction(c *gin.Context)
}

type delivery struct {
	use usecase.Usecase
}

func NewDelivery(u usecase.Usecase) Delivery {
	return &delivery{
		use: u,
	}
}

func (d *delivery) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (dlvr *delivery) NoRoute(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "route not found",
	})
}
