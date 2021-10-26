package root

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// Define remote file
var remotefile = "https://raw.githubusercontent.com/kallefrombosnia/fc_pinger/master/root/regions.json"

func GetRegionsFile() (string, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(cwd, "/root/regions.json"), nil
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
