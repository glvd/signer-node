package ipfs

import (
	"fmt"
	"os"
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
	var bin = "ipfs"
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
