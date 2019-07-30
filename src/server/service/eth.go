package service

import (
	"accelerateNode/src/eth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddPeer ...
func AddPeer(c *gin.Context) {
	urls := c.PostFormArray("urls")
	fmt.Println("post params is", urls)
	eth.AddPeer(urls)
	c.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}

// Peers ...
func Peers(c *gin.Context) {
	resp, err := eth.Peers()
	c.PureJSON(http.StatusOK, gin.H{
		"resp": resp,
		"err":  err,
	})
}

// NodeInfo ...
func NodeInfo(c *gin.Context) {
	resp, err := eth.NodeInfo()
	c.PureJSON(http.StatusOK, gin.H{
		"resp": resp,
		"err":  err,
	})
}
