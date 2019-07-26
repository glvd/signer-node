package sync

import (
	"fmt"
	"signerNode/src/eth"
)

type nodeSync struct {
}

func EthNodeSync() {
	cl := eth.ContractLoader()
	contract := cl.AccelerateNode("0x1d68bc48d6898bad85cf9d2d6cc062d7677543b3")
	peers, err := eth.Peers()
	fmt.Println("peers", peers, err)
	fmt.Println("contract", contract)
}

func IpfsSync() {

}
