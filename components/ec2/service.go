package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Client struct
type Client struct {
	svc *ec2.EC2
}

// New getClient
func New(sess *session.Session) (*Client, error) {
	svc := &Client{
		ec2.New(sess),
	}
	return svc, nil
}
