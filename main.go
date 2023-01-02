package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func printMessage() {
	defer wg.Done()
	fmt.Println(msg)
}

func updateMessage(s string) {
	msg = s
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	wg := sync.WaitGroup{}

	updates := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, word!",
	}

	wg.Add(len(updates))

	for _, update := range updates {
		go updateMessage(update)
		wg.Wait()
	}

	fmt.Println("Done!")
}
