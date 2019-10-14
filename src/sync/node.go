package sync

import (
	"fmt"
	"net"
	"strings"

	"signerNode/src/eth"
	"signerNode/src/ipfs"
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
	node := nodeInfo.Enode

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
	// get accessible nodes
	accessibleNodes := getAccessibleEthNodes(activePeers)
	// sync nodes
	newSignerNodes := difference([]string{node}, cNodes)
	newAccNodes := difference(accessibleNodes, cPeers)

	// node to be deleted
	deleteNodes := difference(cPeers, accessibleNodes)
	var deleteIdx []int
	for _, dNode := range deleteNodes {
		for idx, cNode := range cPeers {
			if cNode == dNode {
				deleteIdx = append(deleteIdx, idx)
			}
		}
	}

	fmt.Println("[待删除节点]", deleteIdx, deleteNodes)

	if len(newAccNodes) > 0 {
		var err error
		for _, n := range newAccNodes {
			_, err = ac.AddEthNodes(auth, []string{n})
		}
		if err != nil {
			fmt.Println("[添加节点失败]", err.Error())
		} else {
			fmt.Println("[添加节点成功]")
		}
		// update gateway info
		for _, node := range newAccNodes {
			ProxySync(node)
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
	// delete rest node
	if len(deleteIdx) > 0 {
		var err error
		for _, n := range deleteIdx {
			_, err = ac.DeleteEthNodes(auth, uint32(n))
		}
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
	// TODO: node priority design
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
		// check if ipfs node in public net
		ip, port := getAddressInfo(peer.Addr)
		conn, err := net.Dial("tcp", ip+":"+port)
		// p2p proxy node
		if err != nil {
			ipfsAddr := "/ipfs/" + nodeID + "/p2p-circuit/ipfs/" + peer.Peer
			peers = append(peers, ipfsAddr)
			// public node
		} else {
			ipfsAddr := peer.Addr + "/ipfs/" + peer.Peer
			publicNodes = append(publicNodes, ipfsAddr)
			conn.Close()
		}
	}
	fmt.Println("[当前IPFS总节点数]", len(peers)+len(publicNodes))
	fmt.Println("[当前IPFS公网节点数]", len(publicNodes))
	// get nodes info
	if len(peers) == 0 {
		fmt.Println("<IPFS节点状态已是最新>")
		return
	}
	cl := eth.ContractLoader()
	ac, auth, client := cl.AccelerateNode()
	defer client.Close()

	cPeers, err := ac.GetIpfsNodes(nil)
	cNodes, err := ac.GetPublicIpfsNodes(nil)
	for _, n := range difference(peers, cPeers) {
		_, err = ac.AddIpfsNodes(auth, []string{n})
	}
	for _, n := range difference(publicNodes, cNodes) {
		_, err = ac.AddPublicIpfsNodes(auth, []string{n})
	}

	if err != nil {
		fmt.Println("[添加节点失败]", err.Error())
	} else {
		fmt.Println("[添加节点成功] ")
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
		url := strs[1]
		ip := strings.Split(url, ":")[0]
		fmt.Println("[trying dial ip]", ip+":30303")
		conn, err := net.Dial("tcp", ip+":30303")
		if err == nil {
			addr := strs[0] + "@" + ip + ":30303"
			accessible = append(accessible, addr)
			conn.Close()
		} else {
			fmt.Println("[dial err]", err)
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
