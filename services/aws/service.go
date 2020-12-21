package aws

import (
	"deploy-helper/components/aws/alb"
	"deploy-helper/components/aws/autoscaling"
	"deploy-helper/components/aws/ec2"
	"deploy-helper/components/aws/ec2metadata"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// IfaceServiceAws AWS service
//go:generate mockery --name IfaceServiceAws --output ../mocks
type IfaceServiceAws interface {
	DeregisterInstanceSelf() error
	RegisterInstanceSelf() error
}

// Client struct
type Client struct {
	ec2metadataSvc ec2metadata.IfaceEc2metadata
	autoscalingSvc autoscaling.IfaceAutoscaling
	ec2Svc         ec2.IfaceEc2
	albSvc         alb.IfaceAlb
}

// New get client
func New() (*Client, error) {
	metadata, err := ec2metadata.New()
	if err != nil {
		return nil, err
	}

	region, err := metadata.GetRegion()
	sess := session.New(&aws.Config{Region: aws.String(region)})

	ec2Svc, err := ec2.New(sess)
	if err != nil {
		return nil, err
	}

	albSvc, err := alb.New(sess)
	if err != nil {
		return nil, err
	}

	autoscalingSvc, err := autoscaling.New(sess)
	if err != nil {
		return nil, err
	}

	return &Client{
		ec2metadataSvc: metadata,
		autoscalingSvc: autoscalingSvc,
		ec2Svc:         ec2Svc,
		albSvc:         albSvc,
	}, nil
}
