package routers

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Router_T interface {
	NewRouter() (Router)
	Start(wg *sync.WaitGroup, port int) ()
}


type Router struct {
	gin.Engine
}

func NewRouter() Router {
	dr := gin.Default()
	dr.SetTrustedProxies([]string{"192.168.0.0/16"})
	return Router{*dr}
}

