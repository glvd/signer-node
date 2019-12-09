package sync

import (
	"testing"
)

func Test_GetRemoteDNSRecords(t *testing.T) {
	arr, err := getRemoteDNSRecords()
	if err != nil {
		t.Log("err", err)
	} else {

		t.Log("arr", arr)
	}
}
