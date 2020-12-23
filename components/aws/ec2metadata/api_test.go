package ec2metadata

import (
	"deploy-helper/components/aws/mocks"
	"fmt"
	"reflect"
	"testing"
)

func TestGetInstanceID(t *testing.T) {
	cases := []struct {
		name               string
		GetMetadataIn      string
		GetMetadataExpResp string
		GetMetadataExpErr  error
		expResp            string
		expErr             error
	}{
		{
			"test",
			"instance-id",
			"instanceID",
			nil,
			"instanceID",
			nil,
		},
		{
			"test",
			"instance-id",
			"",
			fmt.Errorf("get instance id fail"),
			"",
			fmt.Errorf("get instance id fail"),
		},
	}
	for _, c := range cases {
		ec2metadata := new(mocks.Iface3PartyEC2Metadata)
		svc := Client{
			ec2metadata,
		}
		ec2metadata.On("GetMetadata", c.GetMetadataIn).Return(c.GetMetadataExpResp, c.GetMetadataExpErr)
		resp, err := svc.GetInstanceID()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
		if !reflect.DeepEqual(resp, c.expResp) {
			t.Errorf("expected: %v, got: %v", c.expResp, err)
		}
	}
}

func TestGetRegion(t *testing.T) {
	cases := []struct {
		name               string
		GetMetadataIn      string
		GetMetadataExpResp string
		GetMetadataExpErr  error
		expResp            string
		expErr             error
	}{
		{
			"test",
			"placement/region",
			"ap1",
			nil,
			"ap1",
			nil,
		},
		{
			"test",
			"placement/region",
			"",
			fmt.Errorf("get region fail"),
			"",
			fmt.Errorf("get region fail"),
		},
	}
	for _, c := range cases {
		ec2metadata := new(mocks.Iface3PartyEC2Metadata)
		svc := Client{
			ec2metadata,
		}
		ec2metadata.On("GetMetadata", c.GetMetadataIn).Return(c.GetMetadataExpResp, c.GetMetadataExpErr)
		resp, err := svc.GetRegion()
		if !reflect.DeepEqual(err, c.expErr) {
			t.Errorf("expected: %v, got: %v", c.expErr, err)
		}
		if !reflect.DeepEqual(resp, c.expResp) {
			t.Errorf("expected: %v, got: %v", c.expResp, err)
		}
	}
}
