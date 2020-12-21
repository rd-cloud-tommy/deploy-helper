package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// IfaceEc2 interface
//go:generate mockery --name IfaceEc2 --output ../mocks
type IfaceEc2 interface {
	GetTagValue(instanceID string, name string) (string, error)
}

// Client struct
type Client struct {
	svc ec2iface.EC2API
}

// New getClient
func New(sess *session.Session) (*Client, error) {
	svc := &Client{
		ec2.New(sess),
	}
	return svc, nil
}
