package sync

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

const gatewayAPI = "http://gate.dhash.app:8001/upstreams/ETH/targets"

// DNSSync ...
func DNSSync(remoteNodes []string) {
	fmt.Println("<更新网关数据中...>", remoteNodes)
	defer fmt.Println("<更新网关数据完成...>")
	// build node ips
	var ips []string
	for _, node := range remoteNodes {
		if !strings.Contains(node, "enode") {
			continue
		}
		// get ip address
		uri := strings.Split(node, "@")[1]
		ip := strings.Split(uri, ":")[0]
		ips = append(ips, ip)
	}
	if len(ips) == 0 {
		return
	}

	// aws service
	var records []*route53.ResourceRecord
	client := route53.New(session.New())
	// recordSetsInput := &route53.ListResourceRecordSetsInput{
	// 	HostedZoneId:    aws.String("ZULHCRKQ5YDLC"),
	// 	StartRecordName: aws.String("gate.dhash.app"),
	// }
	// sets, err := client.ListResourceRecordSets(recordSetsInput)
	// if err != nil {
	// 	fmt.Println("[get record sets failed]", err.Error())
	// 	return
	// }
	// if len(sets.ResourceRecordSets) != 0 {
	// 	for _, record := range sets.ResourceRecordSets[0].ResourceRecords {
	// 		ips = append(ips, *record.Value)
	// 	}
	// }
	// remove duplicate ip
	ips = removeDuplicateElement(ips)
	for _, ip := range ips {
		awsRecord := &route53.ResourceRecord{Value: aws.String(ip)}
		records = append(records, awsRecord)
	}
	changeInput := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name:            aws.String("gate.dhash.app"),
						ResourceRecords: records,
						TTL:             aws.Int64(60),
						Type:            aws.String("A"),
					},
				},
			},
			Comment: aws.String("gateway"),
		},
		HostedZoneId: aws.String("ZULHCRKQ5YDLC"),
	}
	res, err := client.ChangeResourceRecordSets(changeInput)
	// create health check
	if err != nil {
		fmt.Println("[add resource record fail]", err.Error())
	} else {
		fmt.Println("[add resource record success]", res.String())
	}
	return
}

func createDNSHealthCheck(ip string) {

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
