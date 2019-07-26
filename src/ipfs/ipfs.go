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
	general.RunCMD(i.binPath, "bootstrap", "rm", "all")
	general.RunCMD(i.binPath, "daemon", "--init", "--migrate=true")
}
