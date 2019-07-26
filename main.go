package main

import (
	"signerNode/src/eth"
	"signerNode/src/ipfs"
	"signerNode/assets"
	"signerNode/src/sync"
	"os"
	"path/filepath"
)

func main() {
	// uncompress assets
	restoreAssets()
	// start eth node
	ethWorker := eth.NewEth()
	ethWorker.Init()
	go ethWorker.Start()

	// start ipfs client
	ipfsClient := ipfs.NewIpfs()
	go ipfsClient.Start()
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
