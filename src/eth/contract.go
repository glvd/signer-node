package eth

import (
	"log"

	AccelerateNode "signerNode/src/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type contract struct {
}

// Contracter ...
type Contracter interface {
	AccelerateNode(address string) *AccelerateNode.AccelerateNode
}

// ContractLoader ...
func ContractLoader() Contracter {
	return &contract{}
}

// contract: AccelerateNode init acceleratenode contract
func (c contract) AccelerateNode(contractAddress string) *AccelerateNode.AccelerateNode {
	// TODO
	// gateway redirect to private chain
	client, err := ethclient.Dial("http://139.196.215.224:30303")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(contractAddress)
	instance, err := AccelerateNode.NewAccelerateNode(address, client)
	if err != nil {
		log.Fatal(err)
	}
	client.Close()
	return instance
}
