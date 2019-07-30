package router

import (
	"accelerateNode/src/server/service"

	"github.com/gin-gonic/gin"
)

// IpfsRouter api for ipfs
func IpfsRouter(r *gin.Engine) {
	r.POST("/ipfs/peers", service.AddSwarmPeer)
	r.GET("/ipfs/peers", service.SwarmPeers)
	r.POST("/ipfs/pin", service.PinAdd)
}
