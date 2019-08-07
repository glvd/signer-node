package ipfs

import (
	"context"
	"fmt"
	"time"

	shell "github.com/go-ipfs-restapi-master"
)

// SwarmConnect add swarm peers
func SwarmConnect(address string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sh := shell.NewShell("localhost:5001")
	err := sh.SwarmConnect(ctx, address)
	fmt.Println("connect peer err", err)
	return err
}

// SwarmPeers get swarm peers
func SwarmPeers() *shell.SwarmConnInfos {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sh := shell.NewShell("localhost:5001")
	info, _ := sh.SwarmPeers(ctx)
	return info
}

// BootstrapRmAll remove all bootstrap nodes
func BootstrapRmAll() ([]string, error) {
	sh := shell.NewShell("localhost:5001")
	resp, err := sh.BootstrapRmAll()
	return resp, err
}

// ID get self node info
func ID() (*shell.IDOutput, error) {
	sh := shell.NewShell("localhost:5001")
	resp, err := sh.ID()
	return resp, err
}

// PinAdd ...
func PinAdd(hash string) error {
	sh := shell.NewShell("localhost:5001")
	err := sh.Pin(hash)
	return err
}

// PinLs ...
func PinLs() (map[string]shell.PinInfo, error) {
	sh := shell.NewShell("localhost:5001")
	resp, err := sh.Pins()
	return resp, err
}
