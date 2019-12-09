package sync

import (
	"fmt"
	"signerNode/src/general"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

const (
	HostedZoneID = "ZULHCRKQ5YDLC"
	RecordName   = "gate.dhash.app"
)

// DNSSync ...
func DNSSync(remoteNodes []string) {
	fmt.Println("<更新网关数据中...>", remoteNodes)
	defer fmt.Println("<更新网关数据完成...>")
	var records []string

	// build node records
	for _, node := range remoteNodes {
		if !strings.Contains(node, "enode") {
			continue
		}
		// get ip address
		uri := strings.Split(node, "@")[1]
		ip := strings.Split(uri, ":")[0]
		records = append(records, ip)
	}
	if len(records) == 0 {
		return
	}

	// get remote dns record
	remoteRecords, err := getRemoteDNSRecords()
	if err != nil {
		fmt.Println("<访问远端网关失败> ", err.Error())
		return
	}
	// remove duplicate ip
	ipAdd := removeDuplicateElement(general.DiffStrArray(records, remoteRecords))
	ipDelete := removeDuplicateElement(general.DiffStrArray(remoteNodes, records))
	addDNSRecords(ipAdd)
	deleteDNSRecords(ipDelete)
	return
}

func addDNSRecords(records []string) {
	var awsRecords []*route53.ResourceRecord
	client := route53.New(session.New())

	for _, ip := range records {
		awsRecord := &route53.ResourceRecord{Value: aws.String(ip)}
		awsRecords = append(awsRecords, awsRecord)
	}
	changeInput := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name:            aws.String(RecordName),
						ResourceRecords: awsRecords,
						TTL:             aws.Int64(60),
						Type:            aws.String("A"),
					},
				},
			},
			Comment: aws.String("gateway"),
		},
		HostedZoneId: aws.String(HostedZoneID),
	}
	res, err := client.ChangeResourceRecordSets(changeInput)
	if err != nil {
		return
	}
	fmt.Println("<添加网关记录成功>", res)
}

func deleteDNSRecords(records []string) {
	var awsRecords []*route53.ResourceRecord
	client := route53.New(session.New())

	for _, ip := range records {
		awsRecord := &route53.ResourceRecord{Value: aws.String(ip)}
		awsRecords = append(awsRecords, awsRecord)
	}
	changeInput := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("DELETE"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name:            aws.String(RecordName),
						ResourceRecords: awsRecords,
						Type:            aws.String("A"),
					},
				},
			},
			Comment: aws.String("gateway"),
		},
		HostedZoneId: aws.String(HostedZoneID),
	}
	res, err := client.ChangeResourceRecordSets(changeInput)
	if err != nil {
		return
	}
	fmt.Println("<删除网关记录成功>", res)
}

func getRemoteDNSRecords() ([]string, error) {
	var remoteRecords []string
	client := route53.New(session.New())
	recordSetsInput := &route53.ListResourceRecordSetsInput{
		HostedZoneId:    aws.String(HostedZoneID),
		StartRecordName: aws.String(RecordName),
	}
	sets, err := client.ListResourceRecordSets(recordSetsInput)
	if err != nil {
		fmt.Println("[get record sets failed]", err.Error())
		return remoteRecords, err
	}
	if len(sets.ResourceRecordSets) != 0 {
		for _, record := range sets.ResourceRecordSets[0].ResourceRecords {
			remoteRecords = append(remoteRecords, *record.Value)
		}
	}
	return remoteRecords, nil
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
