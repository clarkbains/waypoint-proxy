package userrouter

import (
	"fmt"
	"sync"
	"github.com/clarkbains/waypoint-proxy/routers"
	"github.com/gin-gonic/gin"
)



func Start(wg *sync.WaitGroup, port int){
	wg.Add(port)
	defer wg.Done()
	ur := routers.NewRouter()
	
	ur.GET("/headers", func(c *gin.Context) {
		c.JSON(200, gin.H{"cwdc-info": c.Request.Header["x-cwdc-login-info"]})
		})
	ur.Run(fmt.Sprintf(":%d", port))
}