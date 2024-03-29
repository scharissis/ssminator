package ssminator

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type Client *ssminator

type ssminator struct {
	awsCfg    aws.Config
	ssmClient *ssm.Client
}

// New returns a pointer to a new instance of `ssminator`.
// To specify an AWS Profile, set the environment variable before constructing your client:
//
//	AWS_PROFILE=profile_name
//
// See: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-profiles
func New() *ssminator {
	cfg, err := newDefaultAWSConfig()
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &ssminator{
		awsCfg:    cfg,
		ssmClient: newDefaultSSMClient(cfg),
	}
}

func newDefaultAWSConfig() (aws.Config, error) {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	return config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
}

func newDefaultSSMClient(cfg aws.Config) *ssm.Client {
	return ssm.NewFromConfig(cfg)
}
