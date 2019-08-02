package sync

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const gatewayAPI = "http://gate.betabb.space:8001/upstreams/ETH/targets"

// ProxySync ...
func ProxySync(remoteNode string) {
	fmt.Println("<更新网关数据中...>")
	data := url.Values{"target": {remoteNode}, "weight": {"15"}}
	body := strings.NewReader(data.Encode())
	resp, err := http.Post(gatewayAPI, "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println("[提交网关失败]", err.Error())
	}
	fmt.Println("[提交网关成功]", resp)
	return
}
