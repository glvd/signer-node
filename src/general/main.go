package general

import (
	"crypto/rand"
	"fmt"
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

// GenerateRandomString return 2 * l long random string
func GenerateRandomString(l int) string {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}
