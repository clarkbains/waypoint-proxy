package userrouter

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http/httputil"
	"sync"

	"github.com/clarkbains/waypoint-proxy/routers"
	"github.com/gin-gonic/gin"
)

type Groups struct {
	GroupId string
	GroupName string
}

type UserInfo struct {
	Id string `json:"id"`
	UserName string `json:"un"`
	Discriminator int `json:"dc"`
	Groups []Groups `json:"gp"`
}
func getUser (s string ) UserInfo {
	if len(s) == 0 {
		s = "eyJpZCI6IjYxODg4ODYxMDY5MDg5MTc5NyIsInNpZCI6W1siOTI5OTU0MTUzNjI5OTA5MDIzIiwid2F5cG9pbnQtcHJveHktdXNlciJdXX0="
	}
	decoded, _ := b64.StdEncoding.DecodeString(s)
	var userData UserInfo
	json.Unmarshal(decoded, &userData)
	return userData
}

func Start(wg *sync.WaitGroup, port int){
	wg.Add(port)
	defer wg.Done()
	ur := routers.NewRouter()
	
	ur.GET("/self", func(c *gin.Context) {
		data := getUser(c.GetHeader("x-cwdc-user"))
		c.String(200, "id: %s, username: %s#%d", data.Id, data.UserName, data.Discriminator)
		})
	ur.Run(fmt.Sprintf(":%d", port))
}