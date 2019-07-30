package service

import (
	"accelerateNode/src/ipfs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PinAdd pin new file
func PinAdd(c *gin.Context) {
	hash := c.PostForm("hash")
	err := ipfs.PinAdd(hash)
	c.JSON(http.StatusOK, gin.H{
		"hash": hash,
		"err":  err.Error(),
	})
}

// PinLs get ipfs pinned list
func PinLs(c *gin.Context) {
	info, err := ipfs.PinLs()
	c.JSON(http.StatusOK, gin.H{
		"list": info,
		"err":  err,
	})
}

// AddSwarmPeer ...
func AddSwarmPeer(c *gin.Context) {
	url := c.PostForm("url")
	fmt.Println("add peer url is", url)
	err := ipfs.SwarmConnect(url)

	c.JSON(http.StatusOK, gin.H{
		"url": url,
		"err": err,
	})
}

// SwarmPeers ...
func SwarmPeers(c *gin.Context) {
	resp := ipfs.SwarmPeers()

	c.PureJSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"peers":  resp.Peers,
	})
}
