package main

import (
	"fmt"

	"github.com/kallefrombosnia/fc_pinger/root"
)

func main() {

	// Print welcome messages
	root.Welcome()

	fmt.Println("Try to load regions...\n\r")

	// Check if we have remotes file
	if !root.CheckForFile() {

		// If doesn exists download file from github
		root.DownloadRegions()

		fmt.Println("Downloading regions.json \n\r")
	}

	root.Init()

}
