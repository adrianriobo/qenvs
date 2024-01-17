package mac

import (
	"fmt"
	"os"
	"strings"

	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"golang.org/x/exp/slices"
)

type HostInformation struct {
	BackedURL   *string
	ProjectName *string
	Region      *string
	Host        *ec2Types.Host
}

// Compose information around dedicated host
func getHostInformation(h ec2Types.Host) HostInformation {
	bi := slices.IndexFunc(
		h.Tags,
		func(t ec2Types.Tag) bool {
			return *t.Key == backedURLTagName
		})
	pi := slices.IndexFunc(
		h.Tags,
		func(t ec2Types.Tag) bool {
			return *t.Key == qenvsContext.ProjectNameTagName
		})
	az := *h.AvailabilityZone
	region := az[:len(az)-1]
	return HostInformation{
		BackedURL:   h.Tags[bi].Value,
		ProjectName: h.Tags[pi].Value,
		Region:      &region,
		Host:        &h,
	}
}

// format for remote backed url when creating the dedicated host
// the backed url from param is used as base and the ID is appended as sub path
func getBackedURL() string {
	if strings.Contains(qenvsContext.BackedURL(), "file://") {
		return qenvsContext.BackedURL()
	}
	return fmt.Sprintf("%s/%s", qenvsContext.BackedURL(), qenvsContext.ID())
}

// Get all dedicated hosts matching the tags + arch
func getMatchingHostsInformation(arch string) ([]HostInformation, error) {
	return getMatchingHostsInStateInformation(arch, nil)
}

// Get all dedicated hosts in available state ordered based on the allocation time
// newer allocations go first
func getMatchingAvailableHostsInformation(arch string) ([]HostInformation, error) {
	as := ec2Types.AllocationStateAvailable
	return getMatchingHostsInStateInformation(arch, &as)
}

// Get all dedicated hosts by tag and state
func getMatchingHostsInStateInformation(arch string, state *ec2Types.AllocationState) ([]HostInformation, error) {
	matchingTags := qenvsContext.GetTags()
	matchingTags[archTagName] = arch
	hosts, err := data.GetDedicatedHosts(data.DedicatedHostResquest{
		Tags: matchingTags,
	})
	if err != nil {
		return nil, err
	}
	var r []HostInformation
	for _, dh := range hosts {
		if state == nil || (state != nil && dh.State == *state) {
			i := slices.IndexFunc(
				dh.Tags,
				func(t ec2Types.Tag) bool {
					return *t.Key == backedURLTagName
				})
			az := *dh.AvailabilityZone
			region := az[:len(az)-1]

			r = append(r, HostInformation{
				BackedURL: dh.Tags[i].Value,
				Region:    &region,
				Host:      &dh,
			})
		}
	}
	// Order by allocation time, first newest
	if len(r) > 1 {
		slices.SortFunc(r, func(a, b HostInformation) int {
			return b.Host.AllocationTime.Compare(*a.Host.AllocationTime)
		})
	}
	return r, nil
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
		az, err = data.GetRandomAvailabilityZone(*r.Region, excludedAZs)
		if err != nil {
			return nil, err
		}
		isOffered, err = data.IsInstanceTypeOfferedByAZ(*r.Region, macTypesByArch[r.Architecture], *az)
		if err != nil {
			return nil, err
		}
		if !isOffered {
			excludedAZs = append(excludedAZs, *az)
		}
	}
	return
}
