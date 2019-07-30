package eth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
func AddPeer(urls []string) (interface{}, error) {
	var result interface{}
	reqBody, err := json.Marshal(map[string]interface{}{
		"method": "admin_addPeer",
		"params": urls,
		"id":     "74",
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("request error is", err)
	}
	json.Unmarshal(body, &result)
	return result, nil
}

// Peers ...
func Peers() ([]string, error) {
	var result Result
	var peers []string

	// parse json
	reqBody, err := json.Marshal(map[string]string{
		"method": "admin_peers",
		"id":     "74",
	})

	if err != nil {
		return peers, err
	}
	// send request
	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return peers, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("request error is", err)
	}

	json.Unmarshal(body, &result)

	for _, peer := range result.Result {
		v, _ := peer.Protocols.(map[string]interface{})
		_, ok := v["eth"].(string)
		if ok == false {
			peers = append(peers, peer.Enode)
		}
	}

	return peers, nil
}

// NodeInfo ...
func NodeInfo() (NodeResult, error) {
	var result NodeResult
	reqBody, err := json.Marshal(map[string]string{
		"method": "admin_nodeInfo",
		"id":     "74",
	})
	if err != nil {
		return result, err
	}
	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	json.Unmarshal(body, &result)
	return result, nil
}
