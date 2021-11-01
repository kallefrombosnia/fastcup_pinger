package main

import (
	"fmt"
	"github.com/kallefrombosnia/fc_pinger/pinger"
)

func main() {

	// Print welcome messages
	Startup()

	fmt.Println("Try to load regions...\n\r")

	// Check if we have remotes file
	if !CheckForFile() {

		// If doesn exists download file from github
		DownloadRegions()

		fmt.Println("Downloading regions.json \n\r")
	}

	// Get all regions from config
	regions := GetRegions()

	// Filter disabled hosts
	var hosts []string
	for _, region := range regions {
		if region.Disabled {
			continue
		}

		hosts = append(hosts, region.Socket)
	}

	// Create new pinger
	p := pinger.NewPinger(hosts)

	// Start pinger
	p.Start()

	fmt.Println("program end")
}
