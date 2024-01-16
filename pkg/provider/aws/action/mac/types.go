package mac

import (
	"github.com/adrianriobo/qenvs/pkg/provider/aws/modules/network"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type MacRequest struct {
	// Prefix for the resources related to mac
	// this is relevant in case of an orchestration with multiple
	// macs on the same stack
	Prefix string

	// Machine params
	Architecture string
	Version      string

	// Location params
	FixedLocation    bool
	Region           *string
	AvailabilityZone *string

	// Topology paras
	Airgap bool
	// For airgap scenario there is an orchestation of
	// a phase with connectivity on the machine (allowing bootstraping)
	// a pahase with connectivyt off where the subnet for the target lost the nat gateway
	airgapPhaseConnectivity network.Connectivity

	// operation control params
	replace bool
	lock    bool

	// This value wil be used to dynamically expand the pool size
	// MaxPoolSize *int

	// dh linkage
	dedicatedHost *ec2Types.Host
}
