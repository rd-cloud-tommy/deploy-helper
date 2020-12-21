package autoscaling

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

type mockedReceiveMsgs struct {
	autoscalingiface.AutoScalingAPI
	DescribeLoadBalancerTargetGroupsResp autoscaling.DescribeLoadBalancerTargetGroupsOutput
	DescribeLoadBalancerTargetGroupsErr  error
	StartInstanceRefreshResp             autoscaling.StartInstanceRefreshOutput
	StartInstanceRefreshErr              error
	UpdateAutoScalingGroupResp           autoscaling.UpdateAutoScalingGroupOutput
	UpdateAutoScalingGroupErr            error
}

func (m mockedReceiveMsgs) DescribeLoadBalancerTargetGroups(in *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	return &m.DescribeLoadBalancerTargetGroupsResp, m.DescribeLoadBalancerTargetGroupsErr
}

func (m mockedReceiveMsgs) StartInstanceRefresh(in *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error) {
	return &m.StartInstanceRefreshResp, m.StartInstanceRefreshErr
}

func (m mockedReceiveMsgs) UpdateAutoScalingGroup(in *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
	return &m.UpdateAutoScalingGroupResp, m.UpdateAutoScalingGroupErr
}

func TestGetTargetGroups(t *testing.T) {
	cases := []struct {
		name                                 string
		autoscalingName                      string
		DescribeLoadBalancerTargetGroupsResp autoscaling.DescribeLoadBalancerTargetGroupsOutput
		DescribeLoadBalancerTargetGroupsErr  error
		expResp                              []TargetGroup
		expErr                               error
	}{
		{
			"test with two tg",
			"autoscaling-foo",
			autoscaling.DescribeLoadBalancerTargetGroupsOutput{
				LoadBalancerTargetGroups: []*autoscaling.LoadBalancerTargetGroupState{
					{
						LoadBalancerTargetGroupARN: aws.String("foo"),
					},
					{
						LoadBalancerTargetGroupARN: aws.String("boo"),
					},
				},
			},
			nil,
			[]TargetGroup{
				{
					aws.String("foo"),
				},
				{
					aws.String("boo"),
				},
			},
			nil,
		},
		{
			"test error",
			"autoscaling-foo",
			autoscaling.DescribeLoadBalancerTargetGroupsOutput{},
			fmt.Errorf("get error"),
			nil,
			fmt.Errorf("get error"),
		},
	}

	for _, c := range cases {
		svc := Client{
			mockedReceiveMsgs{
				DescribeLoadBalancerTargetGroupsResp: c.DescribeLoadBalancerTargetGroupsResp,
				DescribeLoadBalancerTargetGroupsErr:  c.DescribeLoadBalancerTargetGroupsErr,
			},
		}
		resp, err := svc.GetTargetGroups(c.autoscalingName)
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
		if !reflect.DeepEqual(resp, c.expResp) {
			t.Errorf("expected: %v, got: %v", c.expResp, err)
		}
	}
}
