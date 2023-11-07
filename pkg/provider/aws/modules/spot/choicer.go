package spot

import (
	"golang.org/x/exp/slices"

	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const maxScore = 10

func choiceBestSpot(args bestSpotChoiceArgs) (*BestSpotChoice, error) {
	regions, err := data.GetRegions()
	if err != nil {
		return nil, err
	}
	scores, err := placementScores(regions, args.InstanceTypes, 1)
	if err != nil {
		return nil, err
	}
	prices := prices(*args.ProductDescription, regions, args.InstanceTypes)
	// check this
	bestPrice := checkBestOption(prices, scores, describeAvailabilityZones(regions))
	// if bestPrice != nil {
	// 	logging.Debugf("Based on avg prices for instance types %v is az %s, current avg price is %.2f and max price is %.2f with a score of %d",
	// 		args.InstanceTypes, *bestPrice.AvailabilityZone, *bestPrice.AVGPrice, *bestPrice.MaxPrice, bestPrice.Score)
	// }
	return bestPrice, nil
}

func checkBestOption(
	source []bestSpotChoiceState, sps []*ec2.SpotPlacementScore,
	availabilityZones []*ec2.AvailabilityZone) *BestSpotChoice {
	slices.SortFunc(source,
		func(a, b bestSpotChoiceState) int {
			return int(*a.AVGPrice - *b.AVGPrice)
		})
	var score int64 = maxScore
	for score > 3 {
		for _, price := range source {
			idx := slices.IndexFunc(sps, func(item *ec2.SpotPlacementScore) bool {
				// Need transform
				spsZoneName, err := data.GetZoneName(*item.AvailabilityZoneId, availabilityZones)
				if err != nil {
					return false
				}
				return spsZoneName == *price.AvailabilityZone &&
					*item.Score == score
			})
			if idx != -1 {
				price.Region = sps[idx].Region
				price.Score = sps[idx].Score
				return &price
			}
		}
		score--
	}
	return nil
}

func describeAvailabilityZones(regions []string) []*ec2.AvailabilityZone {
	allAvailabilityZones := []*ec2.AvailabilityZone{}
	c := make(chan data.AvailabilityZonesResult)
	for _, region := range regions {
		go data.DescribeAvailabilityZonesAsync(region, c)
	}
	for i := 0; i < len(regions); i++ {
		availabilityZonesResult := <-c
		if availabilityZonesResult.Err == nil {
			allAvailabilityZones = append(allAvailabilityZones, availabilityZonesResult.AvailabilityZones...)
		}
	}
	close(c)
	return allAvailabilityZones
}
