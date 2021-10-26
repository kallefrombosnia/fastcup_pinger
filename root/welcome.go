package root

import (
	"fmt"
)

func Welcome() {

	fmt.Println("=======================================")
	fmt.Println("          FASTCUP.net Pinger           ")
	fmt.Println("=======================================")

	credits()
}

func credits() {
	fmt.Println("Written by kalle\n\r")
}
