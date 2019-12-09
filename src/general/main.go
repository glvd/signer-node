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

// DiffStrArray return the elements in `a` that aren't in `b`.
func DiffStrArray(a, b []string) []string {
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
