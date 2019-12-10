package general

import (
	"testing"
)

func Test_GenerateRandomString(t *testing.T) {
	s := GenerateRandomString(5)
	t.Log("random string", s)
}
