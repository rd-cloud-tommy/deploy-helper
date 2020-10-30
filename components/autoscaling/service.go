package autoscaling

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// Client struct
type Client struct {
	svc *autoscaling.AutoScaling
}

// New getClient
func New(sess *session.Session) (*Client, error) {
	svc := &Client{
		autoscaling.New(sess),
	}
	return svc, nil
}
