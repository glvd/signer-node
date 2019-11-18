package main

import (
	"fmt"
	"log"
	"net"
	"os"
	asset "signerNode/assets"
	"signerNode/src/eth"
	"signerNode/src/general"
	"signerNode/src/ipfs"
	"signerNode/src/sync"
	"time"

	"github.com/jasonlvhit/gocron"
)

var udpstart = true
var (
	AWS_ACCESS_KEY_ID     = "AKIA6EU3XPHQNGX5QRSS"
	AWS_SECRET_ACCESS_KEY = "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"
)

func main() {
	// uncompress assets
	restoreAssets()
	if !general.CheckPortAvailable("30303") || !general.CheckPortAvailable("5001") {
		fmt.Println("[监测到端口被占用,请先结束已启动的ETH/IPFS进程]")
		return
	}
	// start eth node
	ethWorker := eth.NewEth()
	ethWorker.Init()
	go ethWorker.Start()

	// start ipfs client
	ipfsClient := ipfs.NewIpfs()
	go ipfsClient.Start()
	os.Setenv("AWS_ACCESS_KEY_ID", AWS_ACCESS_KEY_ID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", AWS_SECRET_ACCESS_KEY)
	// start sync schedule task
	gocron.Every(60).Seconds().Do(sync.EthNodeSync, ethWorker)
	gocron.Every(60).Seconds().Do(sync.IpfsSync, ipfsClient)
	go gocron.Start()
	// go sync.TokenSync()

	// start udp listening
	fmt.Println("[Listening UDP Port 6067]")
	// listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 6067})
	for {
		if ethWorker.CheckClientReady() {
			break
		}
		time.Sleep(time.Second * 1)
	}
	eth.ListenRemotePinEvent()

	// NAT mapping service
	// for udpstart {
	// 	listenUDP(listener)
	// }
}

func listenUDP(listener *net.UDPConn) {
	fmt.Println("Listening Udp Port...")
	data := make([]byte, 1024)
	_, remoteAddr, err := listener.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}
	log.Printf("<remote node mapping rpc address: %s>\n", remoteAddr.String())
	// sync gateway loadbalance proxy
	// sync.ProxySync(remoteAddr.String())
	return
}

func restoreAssets() {
	dirs := []string{"config", "bin", "node"} // 设置需要释放的目录
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := asset.RestoreAssets("./", dir); err != nil {
			break
		}
	}
}
