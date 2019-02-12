package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/antonosmond/aws-penny-pincher/config"
	"github.com/antonosmond/aws-penny-pincher/ec2"
)

// WaitGroup to keep lambda alive until all go routines have completed
var wg sync.WaitGroup

func handleRequest(ctx context.Context, cfg config.Config) error {

	// process each rule from the config
	for _, rule := range cfg.Rules {
		if err := processRule(&rule); err != nil {
			fmt.Println(err)
		}
	}

	wg.Wait()

	return nil

}

func processRule(rule *config.Rule) error {

	// assume ALL regions if none have been specified
	if len(rule.Regions) == 0 {
		var err error
		rule.Regions, err = ec2.DescribeRegions()
		if err != nil {
			return err
		}
	}

	// loop for each region
	for _, region := range rule.Regions {

		// process each resource type in that region
		for _, resource := range rule.Resources {
			switch resource.Type {
			case "instance":
				if err := processInstances(region, &resource); err != nil {
					fmt.Println(err)
					continue
				}
			default:
				fmt.Printf("%s - skipping unknown resource type: %s\n", region, resource.Type)
			}
		}

	}

	return nil
}
