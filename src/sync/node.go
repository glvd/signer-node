package sync

import (
	"fmt"
	"net"
	"signerNode/src/eth"
	"signerNode/src/ipfs"
	"strings"
)

// EthNodeSync get active nodes add them to the contract
func EthNodeSync() {
	fmt.Println("<同步ETH节点中...>")
	// get active nodes
	activePeers, err := eth.Peers()
	if err != nil {
		fmt.Println("[获取活跃ETH节点失败] ", err.Error())
	} else {
		fmt.Println("[当前活跃ETH节点数] ", len(activePeers))
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

	// add accelerate nodes
	accessibleNodes := getAccessibleEthNodes(newAccNodes)
	if len(accessibleNodes) > 0 {
		ac.AddEthNodes(auth, accessibleNodes)
		_, err := ac.AddEthNodes(auth, accessibleNodes)
		// update gateway info
		for _, node := range accessibleNodes {
			ProxySync(node)
		}
		if err != nil {
			fmt.Println("[添加节点失败]", err.Error())
		} else {
			fmt.Println("[添加节点成功]")
		}
	} else {
		fmt.Println("[已经是最新节点数据]")
	}
	// add signer nodes
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
	// get self node info
	selfInfo, err := ipfs.ID()
	if err != nil {
		fmt.Println("[获取本节点信息失败]", err.Error())
		return
	}
	nodeID := selfInfo.ID
	// get ipfs swarm nodes
	var peers []string
	var publicNodes []string
	resp := ipfs.SwarmPeers()
	for _, peer := range resp.Peers {
		// check if ipfs node is accessible
		ip, port := getAddressInfo(peer.Addr)
		conn, err := net.Dial("tcp", ip+":"+port)
		if err != nil {
			ipfsAddr := "/ipfs/" + nodeID + "/p2p-circuit/ipfs/" + peer.Peer
			peers = append(peers, ipfsAddr)
		} else {
			ipfsAddr := peer.Addr + "/ipfs/" + peer.Peer
			publicNodes = append(publicNodes, ipfsAddr)
			conn.Close()
		}
	}
	fmt.Println("[当前IPFS总节点数]", len(peers)+len(publicNodes))
	fmt.Println("[当前IPFS公网节点数]", len(publicNodes))
	// sync ipfs swarm nodes
	if len(peers) > 0 {
		cl := eth.ContractLoader()
		ac, auth, client := cl.AccelerateNode()
		defer client.Close()

		cPeers, err := ac.GetIpfsNodes(nil)
		cNodes, err := ac.GetPublicIpfsNodes(nil)
		_, err = ac.AddIpfsNodes(auth, difference(peers, cPeers))
		_, err = ac.AddPublicIpfsNodes(auth, difference(publicNodes, cNodes))

		if err != nil {
			fmt.Println("[添加节点失败]", err.Error())
		} else {
			fmt.Println("[添加节点成功] ")
		}
	}
	fmt.Println("<IPFS同步完成>")
	return
}

func getAccessibleEthNodes(addresses []string) []string {
	var accessible []string
	for _, address := range addresses {
		strs := strings.Split(address, "@")
		if len(strs) < 2 {
			continue
		}
		ip := strs[1]
		conn, err := net.Dial("tcp", ip)
		if err == nil {
			accessible = append(accessible, address)
			conn.Close()
		}

	}
	return accessible
}

func getAddressInfo(address string) (ip string, port string) {
	strs := strings.Split(address, "/")
	return strs[2], strs[4]
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
