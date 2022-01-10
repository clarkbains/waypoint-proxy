package userrouter

import (
	"fmt"
	"sync"
	"net/http/httputil"
	"github.com/clarkbains/waypoint-proxy/routers"
	"github.com/gin-gonic/gin"
)



func Start(wg *sync.WaitGroup, port int){
	wg.Add(port)
	defer wg.Done()
	ur := routers.NewRouter()
	
	ur.GET("/dump", func(c *gin.Context) {
		requestDump, _ := httputil.DumpRequest(c.Request, true)

		c.JSON(200, string(requestDump))
		})
	ur.Run(fmt.Sprintf(":%d", port))
}