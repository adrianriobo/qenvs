package mac

import (
	_ "embed"

	"github.com/adrianriobo/qenvs/pkg/provider/aws"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/modules/network"
	"github.com/adrianriobo/qenvs/pkg/util/logging"

	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type MacRequest struct {
	Prefix           string
	Architecture     string
	Version          string
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
	Lock    bool

	// We will create a dh or pick one which will be used to install / replace a mac machine
	dedicatedHost *ec2Types.Host
}

// Create will create a dedicated host and add all tags
// also create the mac machine
func Create(r *MacRequest) (err error) {
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
	// if not only host the mac machine will be created
	if !r.Airgap {
		return r.createMacMachine(dhAZ, host)
	}
	// Airgap scneario requires orchestration
	return r.createAirgapMacMachine(dhAZ, host)
}

// This function will check if mac machine exists based on labeling
// If it will exists It will check if it is locked (where should we add the lock?)
// If lock free we will use replace root volume and set the lock
func Request(r *MacRequest) error {
	hostInformation, err := getMatchingHostsInformation(r.Architecture)
	if err != nil {
		return err
	}
	return replaceMachine(r.Prefix, *hostInformation)
}

func Release(r *MacRequest) error {
	hostInformation, err := getMatchingHostsInformation(r.Architecture)
	if err != nil {
		return err
	}
	return r.releaseLocked(hostInformation)
}

// Initial scenario consider 1 machine
// If we request destroy mac machine it will look for any machine
// and check if it is locked if not locked it will destroy it
func Destroy(prefix, arch string) error {
	hostInformation, err := getMatchingHostsInformation(arch)
	if err != nil {
		return err
	}
	if hostInformation.Host.State == ec2Types.AllocationStateAvailable {
		return aws.DestroyStack(aws.DestroyStackRequest{
			Stackname: stackDedicatedHost,
			Region:    *hostInformation.Region,
			BackedURL: *hostInformation.BackedURL,
		})
	}
	// Dedicated host is not on a valid state to be deleted
	// With same backedURL check if machine is locked
	machineLocked, err := isMachineLocked(prefix, *hostInformation)
	if err != nil {
		return err
	}
	if !machineLocked {
		return aws.DestroyStack(aws.DestroyStackRequest{
			Stackname: stackMacMachine,
			Region:    *hostInformation.Region,
			BackedURL: *hostInformation.BackedURL,
		})
	}
	logging.Debug("nothing to be destroyed")
	return nil
}
