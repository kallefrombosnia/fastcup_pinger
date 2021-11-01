package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

type Region struct {
	Socket   string
	Region   string
	Game     string
	Disabled bool
}

func GetRegions() []Region {

	var regions []Region

	filepath, err := GetRegionsFile()

	if err != nil {
		log.Fatal("Couldn't load regions.json")
	}

	regionsFile, err := os.Open(filepath)

	defer regionsFile.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	jsonParser := json.NewDecoder(regionsFile)

	jsonParser.Decode(&regions)

	return regions
}

// Define remote file
var remotefile = "https://raw.githubusercontent.com/kallefrombosnia/fc_pinger/master/root/regions.json"

func GetRegionsFile() (string, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(cwd, "/pinger/regions.json"), nil
}

func DownloadRegions() {

	filepath, err := GetRegionsFile()

	if err != nil {
		log.Fatal("Cant get filepath.")
	}

	// Create blank file
	file, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}

	// Put content on file
	resp, err := http.Get(remotefile)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	defer file.Close()

}

func CheckForFile() bool {

	filepath, err := GetRegionsFile()

	if err != nil {
		log.Fatal("Cant get filepath.")
	}

	_, err = os.Stat(filepath)

	return !errors.Is(err, os.ErrNotExist)

}
