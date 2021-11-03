package main

import (
	"fmt"

	"github.com/kallefrombosnia/fastcup_pinger/draw"
	"github.com/kallefrombosnia/fastcup_pinger/pinger"
)

func main() {

	// Print welcome messages
	Startup()

	fmt.Println("Try to load regions...\n\r")

	// Check if we have remotes file
	if !CheckForFile() {

		fmt.Println("Downloading regions.json \n\r")

		// If doesn exists download file from github
		DownloadRegions()
	}

	// Get all regions from config
	regions := GetRegions()

	// Filter disabled hosts
	var hosts []string
	var regionsw []string

	for _, region := range regions {
		if region.Disabled {
			continue
		}

		hosts = append(hosts, region.Socket)
		regionsw = append(regionsw, region.RegionName)
	}

	// Create new pinger
	p := pinger.NewPinger(hosts)

	// Start pinger
	p.Start()

	r := draw.NewResults(hosts, regionsw)

	func() {
		for {
			select {
			case msg := <-p.ResponseChan:
				r.CallClear()
				Startup()
				r.ProcessResponse(msg)
				r.PrintResponse()
			}
		}
	}()

}
