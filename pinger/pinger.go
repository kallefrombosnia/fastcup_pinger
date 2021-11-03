package pinger

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Response struct {
	Host    string
	Latency time.Duration
}

type Pinger struct {
	Hosts        []string
	DoneChan     chan bool
	ResponseChan chan Response
}

func NewPinger(hosts []string) Pinger {
	doneChan := make(chan bool)
	respChan := make(chan Response)
	return Pinger{
		Hosts:        hosts,
		DoneChan:     doneChan,
		ResponseChan: respChan,
	}
}

func (p Pinger) ping(host string) {
	u := url.URL{
		Scheme: "wss",
		Host:   host,
		Path:   "/",
	}

	headers := http.Header{
		"sec-gpc": []string{"1"},
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), headers)

	if err != nil {

		fmt.Println("Cant create user host!", err)

		return
	}

	// Init ticker
	go func() {

		ticker := time.NewTicker(5 * time.Second)

		defer ticker.Stop()

		for {
			select {

			case <-p.DoneChan:

				fmt.Println("Done!")

				return

			case <-ticker.C:

				// Latency measure start
				startTime := time.Now()
				err := c.WriteMessage(websocket.TextMessage, []byte(""))

				if err != nil {

					errstring := fmt.Sprintf("%s host: couldnt send an message!", host)

					log.Fatal(errstring)
				}

				_, m, err := c.ReadMessage()
				if err != nil {
					log.Fatal("Cant read message", m)
				}

				// Latency measure end
				endTime := time.Now()

				p.ResponseChan <- Response{
					Host:    host,
					Latency: endTime.Sub(startTime),
				}

			}
		}

	}()
}

func (p Pinger) Start() {

	fmt.Println(p.Hosts)

	for _, host := range p.Hosts {
		p.ping(host)
	}

}
