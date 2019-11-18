package ipfs

import (
	"fmt"
	"os"
	"runtime"
	"signerNode/src/general"
)

type ipfs struct {
	binPath   string
	available bool
}

// Ipfser ...
type Ipfser interface {
	Start()
	CheckClientReady() bool
}

// NewIpfs ...
func NewIpfs() Ipfser {
	var sys = runtime.GOOS
	var bin string
	var err error

	if sys == "windows" {
		bin, err = general.PathParse("/bin/ipfs_win.exe")
	} else if sys == "linux" {
		bin, err = general.PathParse("/bin/ipfs_linux")
		general.RunCMD("chmod", "777", bin)
	} else {
		fmt.Println("os not match")
	}

	if err != nil {
		fmt.Println("err is", err.Error())
	}
	return &ipfs{bin, false}
}

func (i ipfs) Start() {
	dir, _ := general.PathParse("/ipfs_repo")
	// set ipfs path
	os.Setenv("IPFS_PATH", dir)

	general.RunCMD(i.binPath, "init")
	fmt.Println("Starting IPFS Node")
	general.RunCMD(i.binPath, "config", "Swarm.EnableAutoNATService", "--bool", "true")
	general.RunCMD(i.binPath, "config", "Swarm.EnableRelayHop", "--bool", "true")
	general.RunCMD(i.binPath, "daemon", "--routing", "none")
}

// CheckClientReady check client is ready
func (i ipfs) CheckClientReady() bool {
	if !general.CheckPortAvailable("5001") && !general.CheckPortAvailable("4001") {
		return true
	}
	return false
}
