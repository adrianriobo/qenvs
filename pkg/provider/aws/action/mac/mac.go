package mac

import (
	_ "embed"

	"github.com/adrianriobo/qenvs/pkg/provider/aws"
	"github.com/adrianriobo/qenvs/pkg/util/logging"

	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// Create function will create the dedicated host
// it will add several tags to it
//
// * backedURL will be used to create the mac machine stack
// * arch will be used to match new requests for specific arch
// * origin fixed qenvs value
// * instaceTagName id for the qenvs execution
// * (customs) any tag passed as --tags param on the create operation
//
// It will also create a mac machine based on the arch and version setup
// and will set a lock on it

func Create(r *MacRequest) (err error) {
	// Create the dedicated host
	// Region for the dedicated host (must offer the instance type)
	r.Region, err = getRegion(r)
	if err != nil {
		return err
	}
	// Az on the Region to create the dedicated host (must offer the instance type)
	r.AvailabilityZone, err = getAZ(r)
	if err != nil {
		return err
	}
	dh, err := r.createDedicatedHost()
	if err != nil {
		return err
	}
	// Setup the topology and install the mac machine
	if !r.Airgap {
		return r.createMacMachine(*dh)
	}
	return r.createAirgapMacMachine(*dh)
}

// Request could be interpreted as a general way to create / release
//
// Some project will request a mac machine
// based on tags it will check if there is any existing mac machine (based on labels + arch + MaxPoolSize)
//
// if dh machine exists and max pool size has been reached:
//
// if it exists it will check if it is locked
// if no locked it will replace based on version TODO think how this would afferct airgap / proxy mechanism
// it locked it will (wait or return an error)
//
// if dh does not exist or max pool size has not been reched
// create the machine
//
//	...
func Request(r *MacRequest) error {
	hostInformation, err := getMatchingHostsInformation(r.Architecture)
	if err != nil {
		return err
	}
	return r.replaceMachine(*hostInformation)
}

// TODO review how to handle params
// This will release the lock on the machine allowing new request to get the machine
// Currently release will use data to check AMI...as so do not change it
// we are passing those params from cmd
func Release(r *MacRequest) error {
	hostInformation, err := getMatchingHostsInformation(r.Architecture)
	if err != nil {
		return err
	}
	return r.releaseLocked(*hostInformation)
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
