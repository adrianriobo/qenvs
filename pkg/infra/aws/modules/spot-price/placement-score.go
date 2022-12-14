package spotprice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsEC2 "github.com/aws/aws-sdk-go/service/ec2"
	"golang.org/x/exp/slices"
)

// return sps ordered from max to min by score
func GetBestPlacementScores(regions, instanceTypes []string,
	capacity int64) ([]*awsEC2.SpotPlacementScore, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := awsEC2.New(sess)

	sps, err := svc.GetSpotPlacementScores(
		&awsEC2.GetSpotPlacementScoresInput{
			SingleAvailabilityZone: aws.Bool(true),
			InstanceTypes:          aws.StringSlice(instanceTypes),
			RegionNames:            aws.StringSlice(regions),
			TargetCapacity:         aws.Int64(capacity),
			MaxResults:             aws.Int64(maxSpotPlacementScoreResults),
		})
	if err != nil {
		return nil, err
	}
	if len(sps.SpotPlacementScores) == 0 {
		return nil, fmt.Errorf("non available scores")
	}
	slices.SortFunc(sps.SpotPlacementScores,
		func(a, b *awsEC2.SpotPlacementScore) bool {
			return *a.Score > *b.Score
		})
	return sps.SpotPlacementScores, nil
}
