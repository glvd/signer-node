package main

import (
	"signerNode/src/eth"
	"signerNode/src/ipfs"
	"signerNode/assets"
	"signerNode/src/sync"
	"os"
	"fmt"
    "log"
    "net"
	"path/filepath"
	"github.com/jasonlvhit/gocron"
)

var udpstart = true;

func main() {
	// uncompress assets
	restoreAssets()
	// start eth node
	ethWorker := eth.NewEth()
	ethWorker.Init()
	go ethWorker.Start()

	// start sync schedule task
	gocron.Every(30).Seconds().Do(sync.EthNodeSync)
	gocron.Start()
	
	// start udp listening
	fmt.Println("[Listening UDP Port 6067]")
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 6067})
	
	// NAT mapping service
	for udpstart {
		listenUDP(listener)
	}

	// start ipfs client
	ipfsClient := ipfs.NewIpfs()
	ipfsClient.Start()

	// exit
	udpstart = false
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
	sync.ProxySync(remoteAddr.String())
    return;
}

func restoreAssets() {
	dirs := []string{"config", "bin", "node"} // 设置需要释放的目录
	isSuccess := true
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}
}
