package eth

import (
	"fmt"
	"runtime"
	"signerNode/src/general"
)

type eth struct {
	genesisPath string
	nodePath    string
	binPath     string
}

// Ether ...
type Ether interface {
	Init()
	Start()
}

// NewEth ...
func NewEth() Ether {
	var sys = runtime.GOOS
	var bin string
	var err error
	genesis, err := general.PathParse("/config/genesis.json")
	node, err := general.PathParse("/node")

	if sys == "windows" {
		bin, err = general.PathParse("/bin/geth_win.exe")
	} else if sys == "linux" {
		bin, err = general.PathParse("/bin/geth_linux")
	} else {
		fmt.Println("os not match")
	}

	if err != nil {
		fmt.Println("something is wrong", err)
	}

	return &eth{genesis, node, bin}
}

func (e eth) Init() {
	general.RunCMD(e.binPath, "--datadir", e.nodePath, "init", e.genesisPath)
}

func (e eth) Start() {
	general.RunCMD(e.binPath, "--datadir", e.nodePath, "--networkid", "20190723", "--rpcapi", "eth,web3,admin", "--rpc", "--unlock", "\"0\"", "--password", e.nodePath+"/password")
}
