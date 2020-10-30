package alb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Client struct
type Client struct {
	svc *elbv2.ELBV2
}

// New getClient
func New(sess *session.Session) (*Client, error) {
	svc := &Client{
		elbv2.New(sess),
	}
	return svc, nil
}
