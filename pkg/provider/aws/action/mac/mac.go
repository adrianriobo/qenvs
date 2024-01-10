package mac

import (
	_ "embed"
	"fmt"
	"os"

	"golang.org/x/exp/slices"

	"github.com/adrianriobo/qenvs/pkg/manager"
	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/modules/network"
	"github.com/adrianriobo/qenvs/pkg/util/logging"

	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type MacRequest struct {
	Prefix           string
	Architecture     string
	Version          string
	OnlyHost         bool
	OnlyMachine      bool
	FixedLocation    bool
	Region           string
	AvailabilityZone string
	Airgap           bool
	// For airgap scenario there is an orchestation of
	// a phase with connectivity on the machine (allowing bootstraping)
	// a pahase with connectivyt off where the subnet for the target lost the nat gateway
	airgapPhaseConnectivity network.Connectivity

	// If replace flag is true we will handle create / destroy as replace root volume
	Replace bool

	// We will create a dh or pick one which will be used to install / replace a mac machine
	dedicatedHost *ec2Types.Host
}

// this function orchestrate the two stacks related to mac machine
// * the underlaying dedicated host
// * the mac machine
func Create(r *MacRequest) (err error) {
	hosts, err := data.GetDedicatedHosts(data.DedicatedHostResquest{
		Tags: qenvsContext.GetTags(),
	})
	if err != nil {
		return err
	}
	if len(hosts) == 0 {
		// Check if instance type is available on current location
		// region is only needed for dedicated host mac machine got the region from the az
		// of the dedicated host or the request if it creates both at once
		region, err := getRegion(r)
		if err != nil {
			return err
		}
		r.Region = *region
		az, err := getAZ(r)
		if err != nil {
			return err
		}
		r.AvailabilityZone = *az
		// No host id means need to create dedicated host
		host, dhAZ, err := r.createDedicatedHost()
		if err != nil {
			return err
		}
		r.AvailabilityZone = *dhAZ
		r.dedicatedHost = host
	}
	// TODO we need to check wich host we will use
	// some logic based on states for dh + locks on machines on the dh
	r.dedicatedHost = &hosts[0]
	if !r.OnlyHost {
		// if not only host the mac machine will be created
		if !r.Airgap {
			return r.createMacMachine()
		}
		// Airgap scneario requires orchestration
		return r.createAirgapMacMachine()
	}
	return nil
}

// This function will check if mac machine exists based on labeling
// If it will exists It will check if it is locked (where should we add the lock?)
// If lock free we will use replace root volume and set the lock
func Request(r *MacRequest) error {
	hosts, err := data.GetDedicatedHosts(data.DedicatedHostResquest{
		Tags: qenvsContext.GetTags(),
	})
	if err != nil {
		return err
	}
	if len(hosts) == 0 {
		return fmt.Errorf("not hosts found")
	}
	// There is a dedicated host to be used
	// Can be one or n
	// Need to get
	host := hosts[0]
	// Now get the backed url for the host decidedinstance

	i := slices.IndexFunc(host.Tags, func(t ec2Types.Tag) bool { return *t.Key == "instanceID" })
	// dhBackedURL := fmt.Sprintf("%s/%s", qenvsContext.GetBackedURL(), *host.Tags[i].Value)
	dhBackedURL := fmt.Sprintf("%s/%s", qenvsContext.GetBackedURL(), *host.Tags[i].Value)
	logging.Debugf("backedurl %s", dhBackedURL)

	// On the replace we are gonna use the stack for create machine we will set the lock to true
	// then when running the up we will set the target element for the lock
	return nil
}

// TODO add loop until state with target state and timeout?
func CheckState(hostID string) (*string, error) {
	hosts, err := data.GetDedicatedHosts(data.DedicatedHostResquest{
		HostID: hostID,
	})
	if err != nil {
		return nil, err
	}
	if len(hosts) != 1 {
		return nil, fmt.Errorf("unexpected number of hosts")
	}
	state := fmt.Sprintf("%v", &hosts[0].State)
	return &state, nil
}

// Will destroy resources related to machine
func Destroy(r *MacRequest) (err error) {
	var region *string
	if !r.OnlyHost {
		region, err = r.getRegionFromStack(stackMacMachine)
		if err != nil {
			return
		}
		if err = aws.DestroyStackByRegion(*region, stackMacMachine); err != nil {
			return
		}
	}
	if !r.OnlyMachine {
		// We need to get dedicated host region to set on stack
		if region == nil {
			region, err = r.getRegionFromStack(stackDedicatedHost)
			if err != nil {
				return
			}
		}
		return aws.DestroyStackByRegion(*region, stackDedicatedHost)
	}
	return nil
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

// Mac machine can be dinamically moved across regions as it is
// tied to the dedicated host we save the region on the stack to setu[
// the AWS session
func (r *MacRequest) getRegionFromStack(stackName string) (*string, error) {
	stack, err := manager.CheckStack(manager.Stack{
		ProjectName: qenvsContext.GetInstanceName(),
		StackName:   qenvsContext.GetStackInstanceName(stackName),
		BackedURL:   qenvsContext.GetBackedURL()})
	if err != nil {
		return nil, err
	}
	outputs, err := manager.GetOutputs(stack)
	if err != nil {
		return nil, err
	}
	region, ok := outputs[fmt.Sprintf("%s-%s", r.Prefix, outputRegion)].Value.(string)
	if ok {
		return &region, nil
	}
	return nil, fmt.Errorf("%s not found", fmt.Sprintf("%s-%s", r.Prefix, outputRegion))
}
