package autoscaling

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

type targetGroup struct {
	Arn *string
}

// GetTargetGroups get autoscaling target group
func (c *Client) GetTargetGroups(autoscalingName string) ([]targetGroup, error) {
	resp, err := c.svc.DescribeLoadBalancerTargetGroups(&autoscaling.DescribeLoadBalancerTargetGroupsInput{
		AutoScalingGroupName: aws.String(autoscalingName),
	})
	if err != nil {
		return nil, err
	}
	list := []targetGroup{}
	for _, tg := range resp.LoadBalancerTargetGroups {
		t := targetGroup{tg.LoadBalancerTargetGroupARN}
		list = append(list, t)
	}
	return list, nil
}
