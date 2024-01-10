package data

import (
	"context"
	"fmt"
	"sync"

	"github.com/adrianriobo/qenvs/pkg/util"
	"github.com/aws/aws-sdk-go-v2/config"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type DedicatedHostResquest struct {
	HostID string
	Tags   map[string]string
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
	if len(r.Tags) > 0 {
		tagKey := "tag-key"
		dhi.Filter = []ec2Types.Filter{
			{
				Name:   &tagKey,
				Values: maps.Keys(r.Tags)},
		}
	}
	h, err := client.DescribeHosts(context.Background(), dhi)
	if err != nil {
		return nil, err
	}
	if len(h.Hosts) == 0 {
		return nil, fmt.Errorf("dedicated host was not found on current region")
	}
	if len(r.Tags) > 0 {
		return util.ArrayFilter(h.Hosts,
			func(h ec2Types.Host) bool {
				return allTagsMatches(r.Tags, h)
			}), nil
	}
	return h.Hosts, nil
}

// Check if a host contais exactly all tags defined by tags param
func allTagsMatches(tags map[string]string, h ec2Types.Host) bool {
	count := 0
	for k, v := range tags {
		if slices.Contains(h.Tags, ec2Types.Tag{
			Key:   &k,
			Value: &v,
		}) {
			count++
		}
	}
	return count == len(tags)
}
