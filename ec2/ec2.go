package ec2

import (
	"github.com/antonosmond/aws-penny-pincher/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var regions []string
var sess = session.New()

func DescribeRegions() (regions []string, err error) {

	if len(regions) > 0 {
		return regions, nil
	}

	svc := ec2.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	output, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}

	for _, r := range output.Regions {
		regions = append(regions, *r.RegionName)
	}

	return regions, nil

}

func DescribeInstances(region string, filters []config.Filter) (ids []*string, err error) {

	svc := ec2.New(sess, aws.NewConfig().WithRegion(region))

	input := &ec2.DescribeInstancesInput{}
	for _, f := range filters {
		filter := &ec2.Filter{
			Name: aws.String(f.Name),
		}
		for _, v := range f.Values {
			filter.Values = append(filter.Values, &v)
		}
		input.Filters = append(input.Filters, filter)
	}

	err = svc.DescribeInstancesPages(input,
		func(output *ec2.DescribeInstancesOutput, lastPage bool) bool {
			for _, r := range output.Reservations {
				for _, i := range r.Instances {
					ids = append(ids, i.InstanceId)
				}
			}
			return lastPage
		},
	)

	if err != nil {
		return nil, err
	}

	return ids, nil

}

func StopInstances(region string, ids []*string) error {

	svc := ec2.New(sess, aws.NewConfig().WithRegion(region))

	input := &ec2.StopInstancesInput{
		InstanceIds: ids,
	}

	if _, err := svc.StopInstances(input); err != nil {
		return err
	}

	return nil
}
