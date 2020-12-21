package aws

import (
	"deploy-helper/components/aws/autoscaling"
	"deploy-helper/components/aws/mocks"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestDeregisterInstanceSelf(t *testing.T) {
	cases := []struct {
		name                   string
		getInstanceIDExpResp   string
		getInstanceIDExpErr    error
		getTagValueInTag       string
		getTagValueExpResp     string
		getTagValueExpErr      error
		getTargetGroupsExpResp []autoscaling.TargetGroup
		getTargetGroupsExpErr  error
		deregisterExpErr       error
		expErr                 error
	}{
		{
			"test get instance error",
			"",
			fmt.Errorf("get instance error"),
			"",
			"",
			nil,
			[]autoscaling.TargetGroup{},
			nil,
			nil,
			fmt.Errorf("get instance error"),
		},
		{
			"test get tag value error",
			"instanceID",
			nil,
			"aws:autoscaling:groupName",
			"",
			fmt.Errorf("get tag value error"),
			[]autoscaling.TargetGroup{},
			nil,
			nil,
			fmt.Errorf("get tag value error"),
		},
		{
			"test get target groups error",
			"instanceID",
			nil,
			"aws:autoscaling:groupName",
			"asg-svc-foo",
			nil,
			[]autoscaling.TargetGroup{},
			fmt.Errorf("get target groups error"),
			nil,
			fmt.Errorf("get target groups error"),
		},
		{
			"test deregister two target group",
			"instanceID",
			nil,
			"aws:autoscaling:groupName",
			"asg-svc-foo",
			nil,
			[]autoscaling.TargetGroup{
				{Arn: aws.String("1-arn")},
				{Arn: aws.String("2-arn")},
			},
			nil,
			nil,
			nil,
		},
	}

	for _, c := range cases {
		ec2metadata := new(mocks.IfaceEc2metadata)
		autoscalingSvc := new(mocks.IfaceAutoscaling)
		ec2Svc := new(mocks.IfaceEc2)
		albSvc := new(mocks.IfaceAlb)
		svc := &Client{
			ec2metadata,
			autoscalingSvc,
			ec2Svc,
			albSvc,
		}
		ec2metadata.On("GetInstanceID").Return(c.getInstanceIDExpResp, c.getInstanceIDExpErr)
		ec2Svc.On("GetTagValue", c.getInstanceIDExpResp, c.getTagValueInTag).Return(c.getTagValueExpResp, c.getTagValueExpErr)
		autoscalingSvc.On("GetTargetGroups", c.getTagValueExpResp).Return(c.getTargetGroupsExpResp, c.getTargetGroupsExpErr)

		for _, targetGroup := range c.getTargetGroupsExpResp {
			albSvc.On("Deregister", &c.getInstanceIDExpResp, targetGroup.Arn).Return(c.deregisterExpErr)
		}
		err := svc.DeregisterInstanceSelf()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
	}
}

func TestRegisterInstanceSelf(t *testing.T) {
	cases := []struct {
		name                   string
		getInstanceIDExpResp   string
		getInstanceIDExpErr    error
		getTagValueInTag       string
		getTagValueExpResp     string
		getTagValueExpErr      error
		getTargetGroupsExpResp []autoscaling.TargetGroup
		getTargetGroupsExpErr  error
		registerExpErr         error
		expErr                 error
	}{
		{
			"test get instance error",
			"",
			fmt.Errorf("get instance error"),
			"",
			"",
			nil,
			[]autoscaling.TargetGroup{},
			nil,
			nil,
			fmt.Errorf("get instance error"),
		},
		{
			"test get tag value error",
			"instanceID",
			nil,
			"aws:autoscaling:groupName",
			"",
			fmt.Errorf("get tag value error"),
			[]autoscaling.TargetGroup{},
			nil,
			nil,
			fmt.Errorf("get tag value error"),
		},
		{
			"test get target groups error",
			"instanceID",
			nil,
			"aws:autoscaling:groupName",
			"asg-svc-foo",
			nil,
			[]autoscaling.TargetGroup{},
			fmt.Errorf("get target groups error"),
			nil,
			fmt.Errorf("get target groups error"),
		},
		{
			"test deregister two target group",
			"instanceID",
			nil,
			"aws:autoscaling:groupName",
			"asg-svc-foo",
			nil,
			[]autoscaling.TargetGroup{
				{Arn: aws.String("1-arn")},
				{Arn: aws.String("2-arn")},
			},
			nil,
			nil,
			nil,
		},
	}

	for _, c := range cases {
		ec2metadata := new(mocks.IfaceEc2metadata)
		autoscalingSvc := new(mocks.IfaceAutoscaling)
		ec2Svc := new(mocks.IfaceEc2)
		albSvc := new(mocks.IfaceAlb)
		svc := &Client{
			ec2metadata,
			autoscalingSvc,
			ec2Svc,
			albSvc,
		}
		ec2metadata.On("GetInstanceID").Return(c.getInstanceIDExpResp, c.getInstanceIDExpErr)
		ec2Svc.On("GetTagValue", c.getInstanceIDExpResp, c.getTagValueInTag).Return(c.getTagValueExpResp, c.getTagValueExpErr)
		autoscalingSvc.On("GetTargetGroups", c.getTagValueExpResp).Return(c.getTargetGroupsExpResp, c.getTargetGroupsExpErr)

		for _, targetGroup := range c.getTargetGroupsExpResp {
			albSvc.On("Register", &c.getInstanceIDExpResp, targetGroup.Arn).Return(c.registerExpErr)
		}
		err := svc.RegisterInstanceSelf()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
	}
}
