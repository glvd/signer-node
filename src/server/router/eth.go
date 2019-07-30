package router

import (
	"accelerateNode/src/server/service"

	"github.com/gin-gonic/gin"
)

// EthRouter api for eth
func EthRouter(r *gin.Engine) {
	r.GET("/eth/peers", service.Peers)
	r.GET("/eth/nodeinfo", service.NodeInfo)
	r.POST("/eth/peers", service.AddPeer)
}
