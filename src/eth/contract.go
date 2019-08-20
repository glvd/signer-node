package eth

import (
	"log"
	"strings"

	"signerNode/src/contract/accelerateNode"
	"signerNode/src/general"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const keyStore = `{"address":"54c0fa4a3d982656c51fe7dfbdcc21923a7678cb","crypto":{"cipher":"aes-128-ctr","ciphertext":"a8552eef06ae970cd45d27350d6fa910bc9e492d95d41eab980cd8d849cada3b","cipherparams":{"iv":"6e189a9b631710587b704dd62e1fa8c4"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"c57af43e0822709eb04a0d13ce63ca71f147ec40c6ed7a225042a705aa37b16f"},"mac":"b4308156d0b54096d21f51ad831eb18ff9a067e866da9c63685f9ef8ac2483e2"},"id":"fbec8deb-015e-4db2-8a89-43e784618cf4","version":3}`
const nodeAddr = "0x1d68bc48d6898bad85cf9d2d6cc062d7677543b3"

type contract struct {
	keystore string
}

// Contracter ...
type Contracter interface {
	AccelerateNode() (*accelerateNode.AccelerateNode, *bind.TransactOpts, *ethclient.Client)
}

// ContractLoader ...
func ContractLoader() Contracter {
	return &contract{keyStore}
}

// contract: AccelerateNode init acceleratenode contract
func (c contract) AccelerateNode() (*accelerateNode.AccelerateNode, *bind.TransactOpts, *ethclient.Client) {
	// TODO
	auth, err := bind.NewTransactor(strings.NewReader(c.keystore), "123")
	if err != nil {
		log.Fatal(err)
	}
	// gateway redirect to private chain
	node, err := general.PathParse("/node")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(node + "/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(nodeAddr)
	instance, err := accelerateNode.NewAccelerateNode(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance, auth, client
}
