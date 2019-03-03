package main

import (
	"context"
	"io/ioutil"
	"log"
	"sync"

	"github.com/ghodss/yaml"

	"github.com/antonosmond/aws-penny-pincher/config"
	"github.com/antonosmond/aws-penny-pincher/ec2"
)

// WaitGroup to keep lambda alive until all go routines have completed
var wg sync.WaitGroup

func handleRequest(ctx context.Context) error {

	// load config
	log.Println("Loading config...")
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var cfg *config.Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully loaded config!")
	log.Printf("Rules to process: %d\n", len(cfg.Rules))

	// process each rule from the config
	for _, rule := range cfg.Rules {
		if err := processRule(&rule); err != nil {
			log.Println(err)
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
					log.Println(err)
					continue
				}
			default:
				log.Printf("%s - skipping unknown resource type: %s\n", region, resource.Type)
			}
		}

	}

	return nil
}
