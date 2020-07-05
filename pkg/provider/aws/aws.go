package provideraws

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// DiscoverPrice perform price discovery for an array of instance types
func DiscoverPrice(instanceTypes []string, region string) ([]*ec2.SpotPrice, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	var processedInstanceTypes []*string

	for _, t := range instanceTypes {
		processedInstanceTypes = append(processedInstanceTypes, aws.String(t))
	}

	// Create EC2 service client
	ec2Client := ec2.New(sess)

	currentTime := time.Now()

	// Specify the details of the instance that you want to create.
	spotResult, err := ec2Client.DescribeSpotPriceHistory(&ec2.DescribeSpotPriceHistoryInput{
		// todo: get for all available AZs
		AvailabilityZone: aws.String(region + "a"),
		InstanceTypes:    processedInstanceTypes,
		StartTime:        &currentTime,
		EndTime:          &currentTime,
	})

	if err != nil {
		fmt.Println("Could not get spot pricing", err)
		return nil, err
	}

	fmt.Printf("Discovered spot pricing: %v", spotResult.SpotPriceHistory)
	return spotResult.SpotPriceHistory, nil
}
