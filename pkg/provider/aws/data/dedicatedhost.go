package data

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type DedicatedHostResquest struct {
	HostID  string
	TagKeys []string
}

// This function check on all regions for the dedicated host
// and return its state if found or error if no host is found within the hostID
func GetDedicatedHostState(r DedicatedHostResquest) ([]ec2Types.Host, error) {
	regions, err := GetRegions()
	if err != nil {
		return nil, err
	}
	h := make(chan []ec2Types.Host, len(regions))
	e := make(chan string, 1)
	defer close(h)
	defer close(e)
	var wg sync.WaitGroup
	for _, region := range regions {
		wg.Add(1)
		lRegion := region
		go func(h chan []ec2Types.Host) {
			defer wg.Done()
			if hosts, err := getDedicatedHostByRegion(
				r, lRegion); err == nil {
				h <- hosts
			}
		}(h)
	}
	go func(c chan string) {
		wg.Wait()
		c <- "done"
	}(e)
	var hosts []ec2Types.Host
	for {
		exit := false
		select {
		case oHosts := <-h:
			hosts = append(hosts, oHosts...)
		case <-e:
			exit = true
		}
		if exit {
			break
		}
	}
	return hosts, nil
}

// // This funcion check on a specific region if a hosts with the id or the filters exists
// // and returns its state
// // For the time being
// func getDedicatedHostStateByRegion(hostID string, regionName string) (*ec2Types.AllocationState, error) {
// 	hosts, err := getDedicatedHostByRegion(
// 		DedicatedHostResquest{HostID: hostID}, regionName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(hosts) != 1 {
// 		return nil, fmt.Errorf("unexpected number of hosts")
// 	}
// 	return &hosts[0].State, nil
// }

func getDedicatedHostByRegion(r DedicatedHostResquest, regionName string) ([]ec2Types.Host, error) {
	var cfgOpts config.LoadOptionsFunc
	if len(regionName) > 0 {
		cfgOpts = config.WithRegion(regionName)
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), cfgOpts)
	if err != nil {
		return nil, err
	}
	client := ec2.NewFromConfig(cfg)
	// Describe params
	dhi := &ec2.DescribeHostsInput{}
	if len(r.HostID) > 0 {
		dhi.HostIds = []string{r.HostID}
	}
	if len(r.TagKeys) > 0 {
		tagKey := "tag-key"
		dhi.Filter = []ec2Types.Filter{
			{
				Name:   &tagKey,
				Values: r.TagKeys},
		}
	}
	h, err := client.DescribeHosts(context.Background(), dhi)
	if err != nil {
		return nil, err
	}
	if len(h.Hosts) == 0 {
		return nil, fmt.Errorf("dedicated host was not found on current region")
	}
	return h.Hosts, nil
}
