package cmd

import (
	"deploy-helper/components/alb"
	"deploy-helper/components/autoscaling"
	"deploy-helper/components/ec2"
	"deploy-helper/components/ec2metadata"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
)

var allowTraffic = &cobra.Command{
	Use:   "allow-traffic",
	Short: "register instance to target group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start register instance...")

		metadata, err := ec2metadata.New()
		if err != nil {
			panic(err)
		}

		instanceID, err := metadata.GetInstanceID()
		if err != nil {
			panic(err)
		}

		region, err := metadata.GetRegion()
		if err != nil {
			panic(err)
		}

		sess := session.New(&aws.Config{Region: aws.String(region)})

		ec2Svc, err := ec2.New(sess)
		if err != nil {
			panic(err)
		}

		autoscalingName, err := ec2Svc.GetTagValue(instanceID, "aws:autoscaling:groupName")
		if err != nil {
			panic(err)
		}

		autoscalingSvc, err := autoscaling.New(sess)
		if err != nil {
			panic(err)
		}

		targetGroups, err := autoscalingSvc.GetTargetGroups(autoscalingName)
		if err != nil {
			panic(err)
		}

		albSvc, err := alb.New(sess)
		if err != nil {
			panic(err)
		}

		size := len(targetGroups)
		if size > 0 {
			wg := new(sync.WaitGroup)
			wg.Add(size)
			for _, t := range targetGroups {
				go func(targetGroupArn *string) {
					defer wg.Done()
					fmt.Printf("start register instanceID[%s], arn[%s]\n", instanceID, *targetGroupArn)
					err := albSvc.Register(&instanceID, targetGroupArn)
					if err != nil {
						panic(err)
					}
					fmt.Printf("end register instanceID[%s], arn[%s]\n", instanceID, *targetGroupArn)
				}(t.Arn)
			}
			wg.Wait()
		}
	},
}
