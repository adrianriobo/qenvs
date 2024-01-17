package data

import (
	"context"
	"fmt"
	"sync"

	"github.com/adrianriobo/qenvs/pkg/util/logging"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ImageInfo struct {
	Region *string
	Image  *ec2Types.Image
}

// GetAMI on a region based on name and arch
func GetAMI(amiName, amiArch, region *string) (*ImageInfo, error) {
	var cfgOpts config.LoadOptionsFunc
	if len(*region) > 0 {
		cfgOpts = config.WithRegion(*region)
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(), cfgOpts)
	if err != nil {
		return nil, err
	}
	client := ec2.NewFromConfig(cfg)
	var filterName = "name"
	filters := []ec2Types.Filter{
		{
			Name:   &filterName,
			Values: []string{*amiName}}}
	if amiArch != nil {
		var filterArch = "architecture"
		filters = append(filters, ec2Types.Filter{
			Name:   &filterArch,
			Values: []string{*amiArch}})
	}
	result, err := client.DescribeImages(
		context.Background(),
		&ec2.DescribeImagesInput{
			Filters: filters})
	if err != nil {
		logging.Debugf("error checking %s in %s error is %v", *amiName, *region, err)
		return nil, err
	}
	if result == nil || len(result.Images) == 0 {
		logging.Debugf("result len 0 checking %s in %s", *amiName, *region)
		return nil, nil
	}
	logging.Debugf("len %d checking %s in %s", len(result.Images), *amiName, *region)
	if err != nil {
		return nil, err
	}
	return &ImageInfo{
			Region: region,
			Image:  &result.Images[0]},
		nil
}

// IsAMIOffered checks if an ami based on its Name is offered on a specific region
func IsAMIOffered(amiName, amiArch, region *string) (bool, *ImageInfo, error) {
	ami, err := GetAMI(amiName, amiArch, region)
	return ami != nil, ami, err
}

// This function check all regions to get the AMI filter by its name
// it will return the first region where the AMI is offered
func FindAMI(amiName, amiArch *string) (*ImageInfo, error) {
	regions, err := GetRegions()
	if err != nil {
		return nil, err
	}
	r := make(chan *ImageInfo, len(regions))
	e := make(chan string, 1)
	defer close(r)
	var wg sync.WaitGroup
	for _, region := range regions {
		wg.Add(1)
		lRegion := region
		go func(r chan *ImageInfo) {
			defer wg.Done()
			if isOffered, i, _ := IsAMIOffered(
				amiName, amiArch, &lRegion); isOffered {
				r <- i
			}
		}(r)
	}
	go func(e chan string) {
		wg.Wait()
		defer close(e)
		e <- "done"
	}(e)
	select {
	case sAMI := <-r:
		return sAMI, nil
	case <-e:
		return nil, fmt.Errorf("not AMI find with name %s on any region", *amiName)
	}
}
