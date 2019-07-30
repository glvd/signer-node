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
	data := url.Values{"target": {remoteNode}, "weight": {"15"}}
	body := strings.NewReader(data.Encode())
	resp, err := http.Post(gatewayAPI, "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println("post gate way err", err.Error())
	}
	fmt.Println("add proxy success", resp)
	return
}
