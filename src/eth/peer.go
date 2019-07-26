package eth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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
func Peers() (interface{}, error) {
	var result interface{}
	// parse json
	reqBody, err := json.Marshal(map[string]string{
		"method": "admin_peers",
		"id":     "74",
	})

	if err != nil {
		return nil, err
	}
	// send request
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

// NodeInfo ...
func NodeInfo() (interface{}, error) {
	var result interface{}
	reqBody, err := json.Marshal(map[string]string{
		"method": "admin_nodeInfo",
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
		return nil, err
	}
	json.Unmarshal(body, &result)
	return result, nil
}
