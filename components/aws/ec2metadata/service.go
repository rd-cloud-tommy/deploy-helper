package ec2metadata

import (
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// IfaceEc2metadata interface
//go:generate mockery --name IfaceEc2metadata --output ../mocks
type IfaceEc2metadata interface {
	GetInstanceID() (string, error)
	GetRegion() (string, error)
}

// Client struct
type Client struct {
	svc Iface3PartyEC2Metadata
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
