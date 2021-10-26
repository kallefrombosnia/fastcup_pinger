package main

import (
	"fmt"

	"github.com/kallefrombosnia/fc_pinger/root"
)

func main() {

	// Print welcome messages
	root.Welcome()

	// Check if we have remotes file
	if !root.CheckForFile() {

		// If doesn exists download file from github
		root.DownloadRegions()

		fmt.Println("Downloading regions.json \n\r")
	}

	fmt.Println("Try to load regions...\n\r")

	regions := root.GetRegions()
	
}
