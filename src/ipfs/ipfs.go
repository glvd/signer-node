package ipfs

import (
	"fmt"
	"runtime"
	"signerNode/src/general"
)

type ipfs struct {
	binPath string
}

// Ipfser ...
type Ipfser interface {
	Start()
}

// NewIpfs ...
func NewIpfs() Ipfser {
	var sys = runtime.GOOS
	var bin string
	var err error

	if sys == "windows" {
		bin, err = general.PathParse("/bin/ipfs_win.exe")
	} else if sys == "linux" {
		bin, err = "ipfs", nil
	} else {
		fmt.Println("os not match")
	}

	if err != nil {
		fmt.Println("err is", err.Error())
	}
	return &ipfs{bin}
}

func (i ipfs) Start() {
	var sys = runtime.GOOS
	dir, _ := general.PathParse("/ipfs_repo")
	if sys == "windows" {
		general.RunCMD("set IPFS_PATH=" + dir)
	} else if sys == "linux" {
		general.RunCMD("export", "IPFS_PATH="+dir)
	}

	// if general.FileExists(dir + "/config") {
	// 	fmt.Println("IPFS Node Already Inited")
	// } else {
	// 	fmt.Println("Initing IPFS Node")
	// 	general.RunCMD(i.binPath, "init")
	// }
	general.RunCMD(i.binPath, "init")
	fmt.Println("Starting IPFS Node")
	general.RunCMD(i.binPath, "config", "Swarm.EnableAutoNATService", "--bool", "true")
	general.RunCMD(i.binPath, "config", "Swarm.EnableRelayHop", "--bool", "true")
	general.RunCMD(i.binPath, "daemon", "--routing", "none")
}
