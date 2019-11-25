package sync

import (
	"fmt"
	"net"
	"strings"
	"time"

	"signerNode/src/eth"
	"signerNode/src/ipfs"

	"github.com/godcong/dhcrypto"
)

var (
	priKey = "-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIIJKQIBAAKCAgEAxlVkL8ePTNkijdxef2onC4ccmnvHx7d4Vym0IuIUqZ2QFNFU" +
		"Yw/IMT31UFIqegeW5TGs9a43cEBLNNgCyTEKsouPW5ZjSo4c8CeONRdIW9GsRQk0" +
		"fK4XzHVarm9pjPMx2Ln4VynIsWj1P3qy6S9NNIow9z0akBvgzr+pRF1HostGz/RU" +
		"snjxg91o6kGe0gEhAuyzQDoumCS+LnRbHeNj6AsGa773dLvQZmFE9955SfMEy8m6" +
		"UFsv9/oWiKhLdHvKIM+HTXenZ2owI6k/hrcmU03Ruap1nQAwYDpjdRY8AN/3hiAk" +
		"EoO3gU/sx+s3MXPe/iIrZAhKGMXPtVDb2ID1f6u3b8K4tL03SSg1yaimfy77i1vG" +
		"TpHznUBQDwWH5Tz8hXzy6kRAD2AGsvwvbi/QzIJGhtMw1KVLXm9uWTUf6e8vrwZ1" +
		"gSnqzrnT3tBqTWoMHkLJ7oLnj46TqKeh9LPrwBsbsws1bDR4WIsXdpwzgkoTWc7L" +
		"UYyYr4b8ACWKV6QBwKtDvATV36bnx5P51zkATLmmr+zKGtyMn0ldlpZH9YwKpW4K" +
		"NfcqbjknbB+bQMNollHUxF/JuVgw2LGbB4sz2byYYZAoFXqNq31kFngi5nPcHi6j" +
		"8DwgoQ001W8ksim45wRvuqJmUKWRyjJheFCswfS/ge6FVTIUOYEybucjih0CAwEA" +
		"AQKCAgEAlvEkB2z0NYNHVfmx/Xx3bMGOVlAAEpIiRwvZKXcwTIo6vm01sRKwxDEo" +
		"QtHVu/uMrq+ot224iXiVBAmlzSLCxnGKUCTbOkF/6pHgG621hxPC7ON9i6ofOJ7T" +
		"vc/S38+yTwPx6bxGHicIByDFisxSELtfWrqpPaXJ6O9azknDnDTilp+X2iBLhpT6" +
		"JNZ+Hct4KTalkSr2jRnhl488TTnirhW99EBpKfFKQLCkgZRScKJAMyw63K8ZibtO" +
		"bQDQND2F7oSir6VxxTW+n1VOoKNAysN96rS9QBiFuKaXTOP4FJ5fTjel3GVcQlDm" +
		"npv37G4H9xdgOIhKhCH/2zlHp7U1oHX+Qsa+GUq/c3uk1Cuw3bnt9ZvtivOfIeYQ" +
		"I7hoz7mPbU0ijpxFKQRnZT/GmZaYWai6EDpapY9ff7X1r59tVvGMSA3NgmZuz+6t" +
		"Kbg0uOZzJhPzD//Vny7NRqr99/f1VSbJ6YvvBLyWl9OeuQ6IPGzvOi8XaICkqIpT" +
		"tiLvKLYu55noxhAzDrjAvOEKYA07THfS67LFZrj8qZE6lF5umnBl1XwAOH9lQhy/" +
		"8fTDgcu5altQMKVfqr9fh/b7GIcXNBgLtfk1Pu7XGLfj3dpYWvZW5xBO7hFAbauk" +
		"jdfUdFl1lLL3qm/QTr6p32RleGTYN0+Tc2FO3EHmTbDJoFzHPe0CggEBAPDJoJ/4" +
		"Ya+NBrFYfNJ7ndxrjsUdUJE3Ee6l2Nv9JOOFshzmwU/l+Ze0yCcK0etnzEHms+32" +
		"BwBePCmc1zYQc1M+7BA1Tqv+sG8Ho3f3nks/blF4V9YSY6MxH3B+ry1W2wHX9Irf" +
		"tQEY6Dfa96xrV0VB0r9qty1f1YOlfEWJvfDwc14pcgnT2U/zIFju8xGBPDgDakiF" +
		"5NAUlpNS0YwxILnM09TXU2oUN7F8njrftiMyiWOOIQF/UTs+f3jSy8fEPlXGa9xs" +
		"eFMlyHrr2U4exvJyO5zmciLOhrGGopwkzdEc+BINcXiPjJiND7yrW+939VJit0Ej" +
		"/9yA463T6iUSceMCggEBANLdImhPK510+LU+qsF/i7tcrZ0GF/E8TDj9/hc0nWzY" +
		"HDWbRji8LFkG8mMORshHJRaCR4ROStWC7dxy2/M9ewTF8S5H3+XgoS/hvZlx/ulG" +
		"umhWpPxEZT5cmatRD1iYzN9ADqSWu3fIrQ70RBw9PQNPronhPS244lAEocjeOqBH" +
		"bTYwUC+ckDvZEThzAt60Idqiscr194LDgFZ4yNJ08E50a5P8TANekKYIEhuKoT0u" +
		"YHc02pO8LlW7jQXUjJSvKyUKmnRlJ3EcNAx6/ce9mrXtG9vMhNuOqcirrllitvLQ" +
		"aJZVYquNTggJYiQuKK+mcfZanu7hrUh2w5dvFK9i0/8CggEAfMG7O6dR1cdYBGM4" +
		"qUXrUN1Zp7+8ksDZxbCgX7sVdd07n8XfuyoI3BWK7s+oXDP3nN2PtGeY0RQCT/03" +
		"dIepeSRM40j7bhoUCDMI+4uMtKg03Hlh6US140P8aij5UqCB8L6Xsaye9+aTyvzk" +
		"/qzPFs84Bn2gUx4oXoFLliv8Ae5TmCIZOAZPviDWTb3gqt0u+kaqttDI8Rb5vXNX" +
		"py99KUd7Kfg2++tlv8w1n4Nxt2Lj1HU7nK7+w5dqLIvrkaGYOpEIKbj5zvrwmN/C" +
		"Q7umkM+nG3A7CtW+7BQ6BHT9Pq+nyJK2jCS0UAYmdTbD95tLvFfxYwrn8rPFQ7dc" +
		"xcB8yQKCAQBD9ng3jITvPBtJN4iL01MzMVzXxnYDD781g0/ZJOE0ircU5BYPBT95" +
		"9k47dQeFV8Dxb04jq6RdCtUlf3O7A27aC/5/PzU//1WUfDrC8UYK4/wC0yJcGKNV" +
		"JT12RSsgECfAMQJHNDn6EpkMv9gQDgDTR2RnFkzEptlylvuaJV5Z+IuPsqS1o82t" +
		"LHprak5bf02GDXgmhX6gC+kaddWsV3p4nvdpfCD32QvgJ6vGarkrYf4/ja6BfV6l" +
		"zUxXu7kP1yGdz7wWld/PihqQhzeyoD70MhcPkeykY2f/wK3yK2nx+xAqnBywVFv5" +
		"JSUXqjT84DXNBEpDjkNunrDN50SQftb7AoIBAQDtXI9wqY3hAjVRg6bkcMsxJthz" +
		"NRNsMnR7AHI5Cej+exU7etYIIj0d38DJryhnk1SQVvz2fI2u9ixC19A0b/5RiVAM" +
		"/iUOA7NI9d1+Hok3s09Y3a9K+PPotXSKMQ2fFTmHjwEPimvT0d7CL5hrHc4q3jNw" +
		"uMy/gh7qpZnQ6vmWo508oiN8+AoQUU14D6AK+vZ3vdxkmX+AehywgZgT0yw66Hfc" +
		"fArL0LBi2hGgUuwaIWgMINcbUUhSiCtj6O0+2zJXdBk1oV+kqs0yN7YjYEq77gPC" +
		"rchLJIcwUeB/hk1VnY2alsQQthOPW4JZkrqSSg2H9OUCMw5Nw1mt0RCl0HdZ" +
		"\n-----END RSA PRIVATE KEY-----"
	pubKey = "-----BEGIN CERTIFICATE-----\n" +
		"MIIFBjCCAu4CCQDSOHS/Tm+8kjANBgkqhkiG9w0BAQsFADBFMQswCQYDVQQGEwJD" +
		"TjETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0" +
		"cyBQdHkgTHRkMB4XDTE5MTEwNzA2MDA0N1oXDTIwMTEwNjA2MDA0N1owRTELMAkG" +
		"A1UEBhMCQ04xEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0" +
		"IFdpZGdpdHMgUHR5IEx0ZDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB" +
		"AMZVZC/Hj0zZIo3cXn9qJwuHHJp7x8e3eFcptCLiFKmdkBTRVGMPyDE99VBSKnoH" +
		"luUxrPWuN3BASzTYAskxCrKLj1uWY0qOHPAnjjUXSFvRrEUJNHyuF8x1Wq5vaYzz" +
		"Mdi5+FcpyLFo9T96sukvTTSKMPc9GpAb4M6/qURdR6LLRs/0VLJ48YPdaOpBntIB" +
		"IQLss0A6Lpgkvi50Wx3jY+gLBmu+93S70GZhRPfeeUnzBMvJulBbL/f6FoioS3R7" +
		"yiDPh013p2dqMCOpP4a3JlNN0bmqdZ0AMGA6Y3UWPADf94YgJBKDt4FP7MfrNzFz" +
		"3v4iK2QIShjFz7VQ29iA9X+rt2/CuLS9N0koNcmopn8u+4tbxk6R851AUA8Fh+U8" +
		"/IV88upEQA9gBrL8L24v0MyCRobTMNSlS15vblk1H+nvL68GdYEp6s65097Qak1q" +
		"DB5Cye6C54+Ok6inofSz68AbG7MLNWw0eFiLF3acM4JKE1nOy1GMmK+G/AAlilek" +
		"AcCrQ7wE1d+m58eT+dc5AEy5pq/syhrcjJ9JXZaWR/WMCqVuCjX3Km45J2wfm0DD" +
		"aJZR1MRfyblYMNixmweLM9m8mGGQKBV6jat9ZBZ4IuZz3B4uo/A8IKENNNVvJLIp" +
		"uOcEb7qiZlClkcoyYXhQrMH0v4HuhVUyFDmBMm7nI4odAgMBAAEwDQYJKoZIhvcN" +
		"AQELBQADggIBAKvvWUlbzPKlpyGtKKmHU3lkSqLsLed/mEXCiYrCKaXqCmU3peSv" +
		"rfpjTnup/q7DNS62kdUZ283f4P3cdHwEhFawrMyi098UTQjM3Zuy7L1J+ZFM9vHo" +
		"dDDYA7PlucIiDjPEtAd1ydQKis9yTsh/lQW8i4YcTtrm0FRVXHxMJnMuZPzJFAoh" +
		"5omINPQ3bTv7HRvWme4V24rGGzPJQAznjjE9NCwo/NbapXTk85Mjtt0CkX+oqiBB" +
		"m5ZVPGrkdTy1/j3TKnXumSyvlWUl6L/hzhwQA5M4OeG61ez1XJhuedEFcwl+io/X" +
		"JDC/8/M0g5kPlt/i+e3coSsCRdZfXZNCYj//lEXnyXUmmrlHYFc9y6n9VOl79eFn" +
		"quJjX4HrffVFjz/xP/7D5fQnMSSj9m+48mJTo9dKsaqU4W4QvwGglbMEC2yL5Yrw" +
		"I/KejxpBCc7cXc4hYGvA/iIy4NimFO/nQ83CEzwHC7FgOWhwyeJqqUZrYwwMGFTe" +
		"FZ3IZSpxwLxi2LUxqD4HCfke0CwXEiyL//Q+18ntvkDzdPB9jahY1A0iV/y9ZS5J" +
		"OLhSWN0ney5s04NkpX/R2n3AfV+Uazioc0L5pfoAFuhGyf6zYdN9XPMrgXeOap6i" +
		"Zw8J8Apa6dpcllxNY3Gna6y3byc8kclUvhRykxqf4u0xQ5TnTRgOe7AL" +
		"\n-----END CERTIFICATE-----"
	dateKey = time.Date(2019, time.November, 11, 10, 20, 10, 300, time.Local)
)

// EthNodeSync get active nodes add them to the contract
func EthNodeSync(ethWorker eth.Ether) {
	fmt.Println("<同步ETH节点中...>")
	if !ethWorker.CheckClientReady() {
		fmt.Println("<waiting for eth ready>")
		return
	}
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

	// get decoded contract nodes
	cPeers, err := ac.GetEthNodes(nil)
	if err != nil {
		fmt.Println("[获取合约节点失败]", err.Error())
	} else {
		fmt.Println("[合约已有节点数]", len(cPeers))
	}
	cPeers = decodeNodes(cPeers)
	// get decoded contract signer nodes
	cNodes, err := ac.GetSignerNodes(nil)
	if err != nil {
		fmt.Println("[获取合约主节点失败] ", err.Error())
	} else {
		fmt.Println("[合约已有主节点数]", len(cNodes))
	}
	cNodes = decodeNodes(cNodes)
	// filter public network accessible nodes
	accessibleNodes := getAccessibleEthNodes(activePeers, "30303")
	// sync nodes
	newSignerNodes := difference([]string{node}, cNodes)
	newAccNodes := difference(accessibleNodes, cPeers)
	// node to be deleted
	// deleteNodes := difference(cPeers, getAccessibleEthNodes(cPeers, "30303"))
	// var deleteIdx []uint32
	// for _, dNode := range deleteNodes {
	// 	for idx, cNode := range cPeers {
	// 		if cNode == dNode {
	// 			deleteIdx = append(deleteIdx, uint32(idx))
	// 		}
	// 	}
	// }

	// crypto node info && add to contract
	if len(newAccNodes) > 0 {
		var err error
		for _, n := range encodeNodes(newAccNodes) {
			_, err = ac.AddEthNodes(auth, []string{n})
		}
		if err != nil {
			fmt.Println("[添加节点失败]", err.Error())
		} else {
			fmt.Println("[添加节点成功]")
		}
		// update gateway info
	} else {
		fmt.Println("[已经是最新节点数据]")
	}

	// add signer nodes
	if len(newSignerNodes) > 0 {

		_, err := ac.AddSignerNodes(auth, encodeNodes(newSignerNodes))
		if err != nil {
			fmt.Println("<添加主节点失败>", err.Error())
		} else {
			fmt.Println("[添加主节点成功]")
		}
	}

	// delete rest node
	// if len(deleteIdx) > 0 {
	// 	var err error
	// 	_, err = ac.DeleteEthNodes(auth, deleteIdx)

	// 	if err != nil {
	// 		fmt.Println("<删除失效节点失败>", err.Error())
	// 	} else {
	// 		fmt.Println("[删除失效节点成功]")
	// 	}
	// }
	DNSSync(difference(accessibleNodes, cNodes))

	fmt.Println("<同步ETH节点完成>")
	return
}

// IpfsSync get swarm peers and add them to the contract
func IpfsSync(ipfsWorker ipfs.Ipfser) {
	// TODO: node priority design
	fmt.Println("<同步IPFS节点中>")
	if !ipfsWorker.CheckClientReady() {
		fmt.Println("<waiting for ipfs ready>")
		return
	}
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
	cPeers = decodeNodes(cPeers)
	cNodes = decodeNodes(cNodes)

	// add new nodes
	for _, n := range encodeNodes(difference(peers, cPeers)) {
		if n == "" {
			continue
		}
		_, err = ac.AddIpfsNodes(auth, []string{n})
	}
	for _, n := range encodeNodes(difference(publicNodes, cNodes)) {
		if n == "" {
			continue
		}
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

func getAccessibleEthNodes(addresses []string, port string) []string {
	var accessible []string
	for _, address := range addresses {
		strs := strings.Split(address, "@")
		if len(strs) < 2 {
			continue
		}
		url := strs[1]
		ip := strings.Split(url, ":")[0]
		fmt.Println("[trying dial ip]", ip+":"+port)
		conn, err := net.Dial("tcp", ip+":"+port)
		if err == nil {
			addr := strs[0] + "@" + ip + ":" + port
			accessible = append(accessible, addr)
			conn.Close()
		} else {
			fmt.Println("[dial err]", err)
		}

	}
	return accessible
}

func decodeNodes(nodes []string) []string {
	// init contract
	var decodeNodes []string
	decoder := dhcrypto.NewCipherDecode([]byte(priKey), dateKey)

	for _, node := range nodes {
		decoded, err := decoder.Decode(node)
		if err != nil {
			continue
		}
		decodeNodes = append(decodeNodes, string(decoded))
	}
	return decodeNodes
}

func encodeNodes(nodes []string) []string {
	var encodedNodes []string
	encoder := dhcrypto.NewCipherEncoder([]byte(pubKey), 10, time.Now())
	for _, node := range nodes {
		encoded, err := encoder.Encode(node)
		if err != nil {
			continue
		}
		encodedNodes = append(encodedNodes, string(encoded))
	}
	return encodedNodes
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
