package spot

import (
	"strconv"
	"time"

	"golang.org/x/exp/slices"

	"github.com/adrianriobo/qenvs/pkg/util"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsEC2 "github.com/aws/aws-sdk-go/service/ec2"
)

// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSpotPriceHistory.html
var spotQueryFilterProductDescription string = "product-description"

type priceResult struct {
	Prices []bestSpotChoiceState
	Err    error
}

func prices(productDescription string,
	regions, instanceTypes []string) []bestSpotChoiceState {
	worldwidePrices := []bestSpotChoiceState{}
	c := make(chan priceResult)
	for _, region := range regions {
		go spotPriceAsync(
			instanceTypes,
			productDescription,
			region,
			c)
	}
	for i := 0; i < len(regions); i++ {
		spotPriceResult := <-c
		if spotPriceResult.Err == nil {
			worldwidePrices = append(worldwidePrices, spotPriceResult.Prices...)
		}
	}
	close(c)
	return worldwidePrices
}

func spotPriceAsync(instanceTypes []string, productDescription, region string, c chan priceResult) {
	data, err := spotPrice(instanceTypes, productDescription, region)
	c <- priceResult{
		Prices: data,
		Err:    err}

}

func spotPrice(instanceTypes []string, productDescription, region string) (pricesGroup []bestSpotChoiceState, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	if err != nil {
		return nil, err
	}
	svc := awsEC2.New(sess)
	starTime := time.Now().Add(-1 * time.Hour)
	endTime := time.Now()
	history, err := svc.DescribeSpotPriceHistory(
		&awsEC2.DescribeSpotPriceHistoryInput{
			InstanceTypes: aws.StringSlice(instanceTypes),
			Filters: []*awsEC2.Filter{
				{
					Name:   aws.String(spotQueryFilterProductDescription),
					Values: []*string{&productDescription},
				},
			},
			StartTime: &starTime,
			EndTime:   &endTime,
		})
	if err != nil {
		return nil, err
	}
	spotPriceGroups := util.SplitSlice(history.SpotPriceHistory, func(priceData *awsEC2.SpotPrice) bestSpotChoiceState {
		return bestSpotChoiceState{
			AvailabilityZone: priceData.AvailabilityZone,
		}
	})
	logging.Debugf("grouped prices %v", spotPriceGroups)
	for groupInfo, pricesHistory := range spotPriceGroups {
		prices := util.ArrayConvert(pricesHistory, func(priceHisotry *awsEC2.SpotPrice) float64 {
			price, err := strconv.ParseFloat(*priceHisotry.SpotPrice, 64)
			if err != nil {
				// Overcost
				return 100
			}
			return price
		})
		var avg = util.Average(prices)
		groupInfo.AVGPrice = &avg
		slices.SortFunc(prices, func(a, b float64) int { return int(a - b) })
		var max = prices[len(prices)-1]
		groupInfo.MaxPrice = &max
		pricesGroup = append(pricesGroup, groupInfo)
	}
	return
}
