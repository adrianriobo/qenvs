package mac

import (
	"fmt"

	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"golang.org/x/exp/slices"
)

type HostInformation struct {
	BackedURL *string
	Region    *string
	Host      *ec2Types.Host
}

// TODO only handle 1 machine for the time being
func getMatchingHostsInformation() (*HostInformation, error) {
	hosts, err := data.GetDedicatedHosts(data.DedicatedHostResquest{
		Tags: qenvsContext.GetTags(),
	})
	if err != nil {
		return nil, err
	}
	if len(hosts) == 0 {
		return nil, fmt.Errorf("not hosts found")
	}
	i := slices.IndexFunc(
		hosts[0].Tags,
		func(t ec2Types.Tag) bool {
			return *t.Key == backedURLTagName
		})
	az := *hosts[0].AvailabilityZone
	region := az[:len(az)-1]
	return &HostInformation{
		BackedURL: hosts[0].Tags[i].Value,
		Region:    &region,
		Host:      &hosts[0],
	}, nil
}
