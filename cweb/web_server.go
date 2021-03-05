package cweb

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/waittttting/cRPC-client/conf"
)

type WebServer struct {}

func NewWebServer() *WebServer {
	return &WebServer{

	}
}

func (ws *WebServer) Start(config conf.LocalConf)  {

	// todo: 开启特定个数的 goroutine 处理 web 消息

	wh := NewWebHandler()
	r := gin.Default()
	// 遇到跨域问题时，浏览器会先发送 options 请求，进行跨域检查，配置了下面的代码后，会自动进行 option 请求的回复
	// 会回复浏览器是否可以访问，方法的方法（get post）， content-type 等信息
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://google.com"}
	r.Use(cors.New(corsConfig))

	get := r.Group("/middle-platform")
	get.POST("/common", wh.RiskCheck, wh.TokenCheck, wh.Transfer)
	r.Run(fmt.Sprintf(":%v", config.Server.HttpPort))
}
