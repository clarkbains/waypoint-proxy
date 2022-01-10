package routers

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Router_T interface {
	NewRouter() (gin.Engine)
	Start(wg *sync.WaitGroup, port int) ()
}


type Router struct {
	
}

func NewRouter() gin.Engine {
	dr := gin.Default()
	dr.SetTrustedProxies(nil)
	dr.GET("/hr", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hr",
		})})
	//dr.SetTrustedProxies([]string{"192.168.0.0/16"})
	return *dr
}

