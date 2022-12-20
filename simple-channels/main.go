package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		// Accept a string from ping
		s := <-ping
		// Send the string to pong
		pong <- fmt.Sprintf("%s", strings.ToUpper(s))
	}
}

func main() {

    // create 2 channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Waiting for messages...")

    for {
        // print prompt
        fmt.Println("Enter a message:")
        // get user input
        _, _= fmt.Scanln(&s)

        if s ==strings.ToLower("q") {
            break
        }

        ping <- s
        // wait for a response
        response := <-pong
        fmt.Println("You said:", response)
    }
}
