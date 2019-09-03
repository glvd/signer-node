package general

import (
	"fmt"
	"net"
	"os"
)

// CheckPortAvailable ...
func CheckPortAvailable(port string) bool {
	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
		return false
	}
	defer conn.Close()
	fmt.Printf("Port %q is available", port)
	return true
}
