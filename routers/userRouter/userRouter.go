package userrouter
import (
	"github.com/clarkbains/waypoint-proxy/routers"
)

type UserRouter struct {
	routers.Router
}

func (ur UserRouter) Get(){
	ur.GET("/create")
}