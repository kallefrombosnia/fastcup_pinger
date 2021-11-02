package draw

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"time"

	p "github.com/kallefrombosnia/fastcup_pinger/pinger"
)

type Response struct {
	Host    string
	Latency time.Duration
	Region  string
}

type Results struct {
	Hosts []Response
}

func NewResults(hosts []string, regions []string) Results {

	var hostholder []Response

	for key, host := range hosts {

		hostholder = append(hostholder, Response{
			Host:    host,
			Latency: 0,
			Region:  regions[key],
		})
	}

	return Results{
		Hosts: hostholder,
	}

}

func (r Results) ProcessResponse(messageResponse p.Response) {

	for key, host := range r.Hosts {

		if host.Host == messageResponse.Host {

			r.Hosts[key].Latency = messageResponse.Latency

		}
	}

	sort.SliceStable(r.Hosts, func(i, j int) bool {
		return r.Hosts[i].Latency < r.Hosts[j].Latency
	})

}

func (r Results) PrintResponse() {

	for _, result := range r.Hosts {
		fmt.Println(result.Region, " ", result.Latency)
	}

}

// https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (r Results) CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
