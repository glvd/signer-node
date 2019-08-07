package sync

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const gatewayAPI = "http://gate.betabb.space:8001/upstreams/ETH/targets"

// ProxySync ...
func ProxySync(remoteNode string) {
	fmt.Println("<更新网关数据中...>")
	// check if remote node is accessible
	fmt.Println("[远端节点地址]", remoteNode)
	conn, err := net.Dial("tcp", remoteNode)
	if err != nil {
		fmt.Println("[远端节点无响应]", err.Error())
		return
	}
	conn.Close()

	// update gateway proxy
	data := url.Values{"target": {remoteNode}, "weight": {"15"}}
	body := strings.NewReader(data.Encode())
	resp, err := http.Post(gatewayAPI, "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println("[提交网关失败]", err.Error())
	}
	fmt.Println("[提交网关成功]", resp)
	return
}
