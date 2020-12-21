package aws

import (
	"deploy-helper/components/aws/autoscaling"
	"log"
	"sync"
)

// DeregisterInstanceSelf deregister instance self
func (c *Client) DeregisterInstanceSelf() error {
	log.Println("Start deregister self instance...")

	instanceID, err := c.ec2metadataSvc.GetInstanceID()
	if err != nil {
		return err
	}

	targetGroups, err := c.getTargetGroup(instanceID)
	if err != nil {
		return err
	}

	c.deregister(instanceID, targetGroups)
	return nil
}

// RegisterInstanceSelf register instance self
func (c *Client) RegisterInstanceSelf() error {
	log.Println("Start register instance...")

	instanceID, err := c.ec2metadataSvc.GetInstanceID()
	if err != nil {
		return err
	}

	targetGroups, err := c.getTargetGroup(instanceID)
	if err != nil {
		return err
	}

	c.register(instanceID, targetGroups)
	return nil
}

func (c *Client) register(instanceID string, targetGroups []autoscaling.TargetGroup) {
	size := len(targetGroups)
	if size > 0 {
		wg := new(sync.WaitGroup)
		wg.Add(size)
		for _, t := range targetGroups {
			go func(targetGroupArn *string) {
				defer wg.Done()
				log.Printf("Register start with instanceID[%s], arn[%s]\n", instanceID, *targetGroupArn)
				err := c.albSvc.Register(&instanceID, targetGroupArn)
				if err != nil {
					log.Println("Register instance error.", err)
				}
				log.Printf("Register success with instanceID[%s], arn[%s]\n", instanceID, *targetGroupArn)
			}(t.Arn)
		}
		wg.Wait()
	}
}

func (c *Client) deregister(instanceID string, targetGroups []autoscaling.TargetGroup) {
	size := len(targetGroups)
	if size > 0 {
		wg := new(sync.WaitGroup)
		wg.Add(size)
		for _, t := range targetGroups {
			go func(targetGroupArn *string) {
				defer wg.Done()
				log.Printf("Deregister start with instanceID[%s], arn[%s]\n", instanceID, *targetGroupArn)
				err := c.albSvc.Deregister(&instanceID, targetGroupArn)
				if err != nil {
					log.Println("Deregister instance error.", err)
				}
				log.Printf("Deregister success with instanceID[%s], arn[%s]\n", instanceID, *targetGroupArn)
			}(t.Arn)
		}
		wg.Wait()
	}
}

func (c *Client) getTargetGroup(instanceID string) ([]autoscaling.TargetGroup, error) {
	autoscalingName, err := c.ec2Svc.GetTagValue(instanceID, "aws:autoscaling:groupName")
	if err != nil {
		return nil, err
	}

	targetGroups, err := c.autoscalingSvc.GetTargetGroups(autoscalingName)
	if err != nil {
		return nil, err
	}

	return targetGroups, nil
}
