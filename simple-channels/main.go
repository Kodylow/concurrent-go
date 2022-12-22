package main

import (
	"fmt"
	"strings"
)

// shout has two parameters: a receive only chan ping, and a send only chan pong.
// Note the use of <- in function signature. It simply takes whatever
// string it gets from the ping channel, converts it to uppercase and
// appends a few exclamation marks, and then sends the transformed text to the pong channel.
func shout(ping chan string, pong chan string) {
	for {
		// Accept a string from ping
		// the goroutine blocks here until it receives something from the channel
		s, ok := <-ping
        if !ok {
            // returns 0 if returns empty, means channel is closed
        }
		// Send the string to pong
		pong <- fmt.Sprintf("%s", strings.ToUpper(s))
	}
}

func main() {

	// create 2 channels
	ping := make(chan string)
	pong := make(chan string)

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Waiting for messages...")

	for {
		// print prompt
		fmt.Println("Enter a message:")

		// get user input
		_, _ = fmt.Scanln(&s)

		if s == strings.ToLower("q") {
			break
		}

		ping <- s
		// wait for a response
		response := <-pong
		fmt.Println("You said:", response)
	}
}
