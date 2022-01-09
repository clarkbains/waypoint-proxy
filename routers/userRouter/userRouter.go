package userrouter

import (
	"fmt"
	"sync"
	"github.com/clarkbains/waypoint-proxy/routers"
	"github.com/gin-gonic/gin"
)


type UserRouter struct {
	routers.Router
}

func Start(wg *sync.WaitGroup, port int){
	wg.Add(port)
	defer wg.Done()
	ur := UserRouter{routers.NewRouter()}
	ur.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})})
	ur.Run(fmt.Sprintf(":%d", port))
}