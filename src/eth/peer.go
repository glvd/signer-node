package eth

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/rpc"
)

// Network ...
type Network struct {
	Inbound       bool
	LocalAddress  string
	RemoteAddress string
	Static        bool
	Trusted       bool
}

// Peer ...
type Peer struct {
	Caps      []string
	ID        string
	Name      string
	Enode     string
	Network   Network
	Protocols interface{}
}

// Result ...
type Result struct {
	ID      string
	Jsonrpc string
	Result  []Peer
}

// Node self node info
type Node struct {
	ID         string
	Enode      string
	IP         string
	Name       string
	ListenAddr string
	Ports      interface{}
	Protocols  interface{}
}

// NodeResult return node info
type NodeResult struct {
	ID      string
	Jsonrpc string
	Result  Node
}

// AddPeer ...
func AddPeer(urls []string) error {
	reqBody, err := json.Marshal(map[string]interface{}{
		"params": urls,
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := rpc.DialIPC(ctx, endPoint)
	defer client.Close()
	if err != nil {
		return err
	}

	client.Call(nil, "admin_addPeer", reqBody)
	return nil
}

// Peers ...
func Peers() ([]Peer, error) {
	var peers []Peer
	var res []Peer
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := rpc.DialIPC(ctx, endPoint)
	if err != nil {
		return res, err
	}
	defer client.Close()

	client.Call(&peers, "admin_peers")

	for _, peer := range peers {
		v, _ := peer.Protocols.(map[string]interface{})
		_, ok := v["eth"].(string)
		if ok == false {
			res = append(res, peer)
		}
	}
	return res, nil
}

// NodeInfo ...
func NodeInfo() (Node, error) {
	var result Node

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := rpc.DialIPC(ctx, endPoint)
	defer client.Close()
	if err != nil {
		return result, err
	}
	client.Call(&result, "admin_nodeInfo")
	return result, nil
}
