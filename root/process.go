package root

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func Init() {

	regions := GetRegions()

	fmt.Println("Launching client sockets...\n\r")

	//message := make(chan string)
	done := make(chan bool)

	type ResultsData struct {
		Socket string
		Start  int64
		End    int64
	}

	type Result struct {
		Data *ResultsData
	}

	results := make(chan Result)

	for _, socket := range regions {

		if socket.Disabled != false {
			continue
		}

		func(socket Region) {

			u := url.URL{
				Scheme: "wss",
				Host:   socket.Socket,
				Path:   "/",
			}

			headers := http.Header{
				"sec-gpc": []string{"1"},
			}

			c, _, err := websocket.DefaultDialer.Dial(u.String(), headers)

			if err != nil {

				fmt.Println("Cant create user socket!", err)

				return
			}

			// Init ticker
			go func() {

				ticker := time.NewTicker(5 * time.Second)

				defer ticker.Stop()

				for {
					select {

					case <-done:

						fmt.Println("Done!")

						return

					case <-ticker.C:

						err := c.WriteMessage(websocket.TextMessage, []byte(""))

						var resolve ResultsData

						resolve.Socket = socket.Socket
						resolve.Start = time.Now().UTC().UnixNano() / 1e6

						if err != nil {

							errstring := fmt.Sprintf("%s socket: couldnt send an message!", socket.Socket)

							log.Fatal(errstring)
						}

						_, m, err := c.ReadMessage()

						if err != nil {
							log.Fatal("Cant read message", m)
						}

						resolve.End = time.Now().UTC().UnixNano() / 1e6

					}
				}

			}()

			if err != nil {
				log.Fatal("Cant send message")
			}

		}(socket)

	}

	ticker := time.NewTicker(100)

	for {
		select {

		case <-ticker.C:

		case <-results:
			fmt.Println(results)
		}
	}

}
