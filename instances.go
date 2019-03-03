package main

import (
	"log"

	"github.com/antonosmond/aws-penny-pincher/config"
	"github.com/antonosmond/aws-penny-pincher/ec2"
)

func processInstances(region string, resource *config.Resource) error {

	wg.Add(1)

	go func() {

		defer wg.Done()

		ids, err := ec2.DescribeInstances(region, resource.Filters)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s - %d instances\n", region, len(ids))
		if len(ids) == 0 {
			return
		}
		for _, a := range resource.Actions {
			switch a {
			case "stop":
				log.Printf("%s - stopping instances\n", region)
				ec2.StopInstances(region, ids)
			default:
				log.Printf("%s - skipping unknown action: %s\n", region, a)
			}
		}
	}()

	return nil
}
