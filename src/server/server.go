package server

import (
	"accelerateNode/src/server/router"

	"github.com/gin-gonic/gin"
)

// Start ...
func Start() {
	r := gin.Default()
	router.EthRouter(r)
	router.IpfsRouter(r)
	r.Run("0.0.0.0:6066") // listen and serve on 0.0.0.0:6066
}
