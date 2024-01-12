package mac

import (
	"fmt"
	"os"

	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"golang.org/x/exp/slices"
)

type HostInformation struct {
	BackedURL *string
	Region    *string
	Host      *ec2Types.Host
}

// TODO only handle 1 machine for the time being
func getMatchingHostsInformation(arch string) (*HostInformation, error) {
	matchingTags := qenvsContext.GetTags()
	matchingTags[archTagName] = arch
	hosts, err := data.GetDedicatedHosts(data.DedicatedHostResquest{
		Tags: matchingTags,
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

// checks if the machine can be created on the current location (machine type is available on the region)
// if it available it returns the region name
// if not offered and machine should be created on the region it will return an error
// if not offered and machine could be created anywhere it will get a region offering the machine and return its name
func getRegion(r *MacRequest) (*string, error) {
	logging.Debugf("checking if %s is offered at %s",
		r.Architecture,
		os.Getenv("AWS_DEFAULT_REGION"))
	isOffered, err := data.IsInstanceTypeOfferedByRegion(
		macTypesByArch[r.Architecture],
		os.Getenv("AWS_DEFAULT_REGION"))
	if err != nil {
		return nil, err
	}
	if isOffered {
		region := os.Getenv("AWS_DEFAULT_REGION")
		return &region, nil
	}
	if !isOffered && r.FixedLocation {
		return nil, fmt.Errorf("the requested mac %s is not available at the current region %s and the fixed-location flag has been set",
			r.Architecture,
			os.Getenv("AWS_DEFAULT_REGION"))
	}
	// We look for a region offering the type of instance
	return data.LokupRegionOfferingInstanceType(
		macTypesByArch[r.Architecture])
}

// Get a random AZ from the requested region, it ensures the az offers the instance type
func getAZ(r *MacRequest) (az *string, err error) {
	isOffered := false
	var excludedAZs []string
	for !isOffered {
		az, err = data.GetRandomAvailabilityZone(r.Region, excludedAZs)
		if err != nil {
			return nil, err
		}
		isOffered, err = data.IsInstanceTypeOfferedByAZ(r.Region, macTypesByArch[r.Architecture], *az)
		if err != nil {
			return nil, err
		}
		if !isOffered {
			excludedAZs = append(excludedAZs, *az)
		}
	}
	return
}
