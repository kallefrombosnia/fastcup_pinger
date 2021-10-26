package root

import (
	"encoding/json"
	"log"
	"os"
)

type Region struct {
	Socket   string
	Region   string
	Game     string
	Disabled string
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
