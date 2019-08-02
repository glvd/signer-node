package sync

import (
	"fmt"
	"signerNode/src/eth"
	"signerNode/src/ipfs"
)

// EthNodeSync get active nodes add them to the contract
func EthNodeSync() {
	fmt.Println("<同步ETH节点中...>")
	// get active nodes
	activePeers, err := eth.Peers()
	if err != nil {
		fmt.Println("[获取活跃节点失败] ", err.Error())
	} else {
		fmt.Println("[当前活跃节点数] ", len(activePeers))
	}
	// get self node info
	nodeInfo, err := eth.NodeInfo()
	if err != nil {
		fmt.Println("[获取本节点信息失败] ", err.Error())
	}
	node := nodeInfo.Result.Enode

	// init contract
	cl := eth.ContractLoader()
	ac, auth, client := cl.AccelerateNode()
	defer client.Close()

	// get contract nodes
	cPeers, err := ac.GetEthNodes(nil)
	if err != nil {
		fmt.Println("[获取合约节点失败]", err.Error())
	} else {
		fmt.Println("[合约已有节点数]", len(cPeers))
	}

	// get contract signer nodes
	cNodes, err := ac.GetSignerNodes(nil)
	if err != nil {
		fmt.Println("[获取合约主节点失败] ", err.Error())
	} else {
		fmt.Println("[合约已有主节点数]", len(cNodes))
	}

	// sync nodes
	newSignerNodes := difference([]string{node}, cNodes)
	newAccNodes := difference(activePeers, cPeers)
	if len(newAccNodes) > 0 {
		_, err := ac.AddEthNodes(auth, newAccNodes)
		if err != nil {
			fmt.Println("[添加节点失败]", err.Error())
		} else {
			fmt.Println("[添加节点成功]")
		}
	} else {
		fmt.Println("[已经是最新节点数据]")
	}
	if len(newSignerNodes) > 0 {
		_, err := ac.AddSignerNodes(auth, newSignerNodes)
		if err != nil {
			fmt.Println("<添加主节点失败>", err.Error())
		} else {
			fmt.Println("[添加主节点成功]")
		}
	}
	fmt.Println("<同步ETH节点完成>")
	return
}

// IpfsSync get swarm peers and add them to the contract
func IpfsSync() {
	fmt.Println("<同步IPFS节点中>")
	// get ipfs swarm nodes
	var peers []string
	resp := ipfs.SwarmPeers()
	for _, peer := range resp.Peers {
		ipfsAddr := peer.Addr + "/ipfs/" + peer.Peer
		peers = append(peers, ipfsAddr)
	}
	fmt.Println("[当前IPFS节点数]", len(peers))
	// sync ipfs swarm nodes
	if len(peers) > 0 {
		cl := eth.ContractLoader()
		ac, auth, client := cl.AccelerateNode()
		defer client.Close()
		cPeers, err := ac.GetIpfsNodes(nil)
		_, err = ac.AddIpfsNodes(auth, difference(peers, cPeers))
		if err != nil {
			fmt.Println("[添加节点失败]", err.Error())
		} else {
			fmt.Println("[添加节点成功] ")
		}
	}
	fmt.Println("<IPFS同步完成>")
	return
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
