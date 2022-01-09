package routers
import (

	"github.com/gin-gonic/gin"
)
type Router struct {
	gin.RouterGroup
}

func (Router) Get()