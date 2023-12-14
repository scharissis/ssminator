package ssminator

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmtypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

// InstanceInformation is like github.com/aws/aws-sdk-go-v2/service/ssm/types.InstanceInformation
// but without the pointers.
type InstanceInformation struct {
	//ssmtypes.InstanceInformation // copy of the original

	InstanceID   string
	ComputerName string
}

// DescribeInstanceInformation returns SSM's 'DescribeInstanceInformation' Output.
// This is information from all SSM Fleet-Managed instances.
func (s ssminator) DescribeInstanceInformation() ([]InstanceInformation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	params := &ssm.DescribeInstanceInformationInput{}
	out, err := s.ssmClient.DescribeInstanceInformation(ctx, params)
	if err != nil {
		return nil, err
	}
	iiList := convertInstanceInformation(out.InstanceInformationList...)
	if out.NextToken != nil {
		log.Printf("TODO: handle pagination")
	}

	return iiList, err
}

func convertInstanceInformation(instanceInfos ...ssmtypes.InstanceInformation) []InstanceInformation {
	out := make([]InstanceInformation, len(instanceInfos))
	for i, ii := range instanceInfos {
		out[i] = InstanceInformation{
			InstanceID:   aws.ToString(ii.InstanceId),
			ComputerName: aws.ToString(ii.ComputerName),
		}
	}
	return out
}
