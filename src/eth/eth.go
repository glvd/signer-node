package eth

import (
	"signerNode/src/general"
)

type eth struct {
	nodePath  string
	binPath   string
	available bool
}

// Ether ...
type Ether interface {
	CheckClientReady() bool
}

var bin string
var endPoint string

// NewEth ...
func NewEth() Ether {
	var bin = "geth"
	var node = "/root/.ethereum"
	endPoint = "/root/.ethereum/geth.ipc"

	return &eth{node, bin, false}
}

// CheckClientReady check client is ready
func (e eth) CheckClientReady() bool {
	if !general.CheckPortAvailable("30303") {
		return true
	}
	return false
}
