package eth

import (
	"fmt"
	"math/big"
	dhToken "signerNode/src/contract/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ListenRemotePinEvent listen remote node pin success event
func ListenRemotePinEvent() {
	var opts bind.WatchOpts
	eventCh := make(chan *dhToken.DhTokenPinSuccess)

	// init contract
	cl := ContractLoader()
	instance, _, client := cl.DHToken()
	defer client.Close()

	// watch pin event
	subs, err := instance.WatchPinSuccess(&opts, eventCh, []common.Address{})
	if err != nil {
		fmt.Println("Subscribe Failed: ", err)
		return
	}

	// keep listening
	fmt.Println("---------------listening remote pin event----------------")
	for {
		select {
		case err := <-subs.Err():
			fmt.Println("sub error", err) // Error
		case l := <-eventCh:
			res := common.NewMixedcaseAddress(l.User)
			fmt.Println("<--------user-------->", res.Original())
			fmt.Println("<--------hash-------->", l.Hash)
			fmt.Println("<--------node-------->", l.Date)
			cl := ContractLoader()
			dt, auth, client := cl.DHToken()
			defer client.Close()
			dt.Transfer(auth, res.Address(), big.NewInt(1))

		}
	}

}
