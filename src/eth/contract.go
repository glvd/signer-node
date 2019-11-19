package eth

import (
	"log"
	"strings"

	"signerNode/src/contract/accelerateNode"
	dhToken "signerNode/src/contract/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const keyStore = `{"address":"945d35cd4a6549213e8d37feb5d708ec98906902","crypto":{"cipher":"aes-128-ctr","ciphertext":"649f5c7def3f345c39dc6f10e5438e179a5f06ff1d9ef2467ff7c84ec94f1a2a","cipherparams":{"iv":"0d66dfbc2c978ed1989e2fca05c16abe"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"547ed9895deda897adbe09058ebfb24fb5036695d490c2127da45c4f7ec9e4a8"},"mac":"db76804c69ceb8705de1a73ae0caf4761bd73c3d42aa43f801c03e7fdda6adff"},"id":"9aaeec2d-d639-425a-83f7-a0956dcc78a1","version":3}`
const nodeContractAddr = "0xbaEEB7a3AF34a365ACAa1f8464A3374B58ac9889"
const tokenContractAddr = "0x9064322CfeE623A447ba5aF0dA6AD3341c073535"

type contract struct {
	keystore string
}

// Contracter ...
type Contracter interface {
	AccelerateNode() (*accelerateNode.AccelerateNode, *bind.TransactOpts, *ethclient.Client)
	DHToken() (*dhToken.DhToken, *bind.TransactOpts, *ethclient.Client)
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
	client, err := ethclient.Dial("/root/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(nodeContractAddr)
	instance, err := accelerateNode.NewAccelerateNode(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance, auth, client
}

// contract: DHToken init DHToken contract
func (c contract) DHToken() (*dhToken.DhToken, *bind.TransactOpts, *ethclient.Client) {
	// TODO
	auth, err := bind.NewTransactor(strings.NewReader(c.keystore), "123")
	if err != nil {
		log.Fatal(err)
	}
	// gateway redirect to private chain
	client, err := ethclient.Dial("/root/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(tokenContractAddr)
	instance, err := dhToken.NewDhToken(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance, auth, client
}
