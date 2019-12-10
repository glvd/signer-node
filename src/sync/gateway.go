package sync

import (
	"fmt"
	"signerNode/src/aws"
	"signerNode/src/general"
	"strings"
)

const (
	// test
	HostedZoneID = "Z2D12816OG1BTZ"
	RecordName   = "gate.betabb.space"
	// HostedZoneID = "ZULHCRKQ5YDLC"
	// RecordName   = "gate.dhash.app"
)

// DNSSync ...
func DNSSync(nodes []string) {
	defer fmt.Println("<更新网关数据完成...>")
	var records []string
	// build node records
	for _, node := range nodes {
		if !strings.Contains(node, "enode") {
			continue
		}
		// get ip address
		uri := strings.Split(node, "@")[1]
		ip := strings.Split(uri, ":")[0]
		records = append(records, ip)
	}
	fmt.Println("<更新网关数据中...>", records)

	if len(records) == 0 {
		return
	}
	dnsService := aws.NewRoute53()

	// get remote dns record
	var remoteIPs []string
	remoteRecordSets, err := dnsService.GetRecordSets()
	if err != nil {
		fmt.Println("<访问远端网关失败> ", err.Error())
		return
	}
	if len(remoteRecordSets) != 0 {
		for _, recordSet := range remoteRecordSets {
			remoteIPs = append(remoteIPs, *recordSet.ResourceRecords[0].Value)
		}
	}
	// add new record
	ipAdd := removeDuplicateElement(general.DiffStrArray(records, remoteIPs))
	fmt.Println("[resource to be added]", ipAdd)
	setsAdd := dnsService.BuildMultiValueRecordSets(ipAdd)
	if len(setsAdd) > 0 {
		res, err := dnsService.ChangeSets(setsAdd, "UPSERT")
		if err != nil {
			fmt.Println("[add resource record fail]", err.Error())
		} else {
			fmt.Println("[add resource record success]", res.String())
		}
	}

	// delete record out of date
	failedSets := dnsService.FilterFailedRecords(remoteRecordSets)
	fmt.Println("[resource to be deleted]", len(failedSets))
	if len(failedSets) > 0 {
		res, err := dnsService.ChangeSets(failedSets, "DELETE")
		if err != nil {
			fmt.Println("[delete resource record fail]", err.Error())
		} else {
			fmt.Println("[delete resource record success]", res.String())
		}
	}

	return
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
