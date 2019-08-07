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
		bin, err = "geth", nil
	} else {
		fmt.Println("os not match")
	}

	if err != nil {
		fmt.Println("something is wrong", err)
	}

	return &eth{genesis, node, bin}
}

func (e eth) Init() {
	if general.FileExists(e.nodePath + "/geth/nodekey") {
		fmt.Println("ETH Chain Data Already Inited")
	} else {
		general.RunCMD(e.binPath, "--datadir", e.nodePath, "init", e.genesisPath)
	}
}

func (e eth) Start() {
	var sys = runtime.GOOS
	if sys == "windows" {
		general.RunCMD(e.binPath, "--datadir", e.nodePath, "--networkid", "20190723", "--rpc", "--rpcaddr", "0.0.0.0", "--rpccorsdomain", "*", "--rpcapi", "eth,web3,admin,net", "--unlock", "0", "--password", e.nodePath+"/password", "--mine")
	} else {
		general.RunCMD(e.binPath, "--datadir", e.nodePath, "--networkid", "20190723", "--rpc", "--rpcaddr", "0.0.0.0", "--rpccorsdomain", "*", "--rpcapi", "eth,web3,admin,net", "--allow-insecure-unlock", "--unlock", "0", "--password", e.nodePath+"/password", "--mine")
	}
}
