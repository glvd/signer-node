package general

import (
	"net"
)

// CheckPortAvailable ...
func CheckPortAvailable(port string) bool {
	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
