package cweb

import (
	"cx-mp-gateway/crisk"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/waittttting/cRPC-common/cerr"
	"net/http"
)

type WebHandler struct {}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

// 风控
func (wh *WebHandler) RiskCheck(c *gin.Context) {
	if !crisk.GRS.CheckRisk(c) {
		// 风控校验失败
		c.JSON(http.StatusOK, cerr.ErrBusy)
	}
}

// token 检查
func (wh *WebHandler) TokenCheck(c *gin.Context) {

	header, _, info := getHeaderFromCtx(c)
	println(info)
	fmt.Printf("%v", header)
}

// 转发
func (wh *WebHandler) Transfer(c *gin.Context) {

}
