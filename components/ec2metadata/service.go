package ec2metadata

import (
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Client struct
type Client struct {
	svc *ec2metadata.EC2Metadata
}

// New getClient
func New() (*Client, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	metadata := ec2metadata.New(sess)
	svc := &Client{
		metadata,
	}
	return svc, nil
}
