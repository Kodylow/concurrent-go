package main

import (
    "fmt"
    "time"
)

fun listenToChan(ch chan string) {
    for {
        i := <-ch
        fmt.Println("Got:", i)

        time.Sleep(1 * time.Second)
    }
}

func main() {
    // adding the int at the end creates a buffered channel
    // useful when :
    //     1. you know how many goroutines you've launched
    //     2. you want to limit the number of goroutines you launch
    //     3. you want to limit the amount of work that's queued up
    ch := make(chan int, 10)

    go listenToChan(ch)

    for i := 0; i < 10; i++ {
        fmt.Println("Sending:", i)
        ch <- i
        fmt.Println("Sent:", i)
    }

    fmt.Println("Done")
    close(ch)
}