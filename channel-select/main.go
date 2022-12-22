package main

import (
    "fmt"
    "time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
        ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
    for {
        time.Sleep(3 * time.Second)
        ch <- "This is from server 2"
    }}
}

func main() {
    fmt.Println("Select with channels\n-------------------")

    ch1 := make(chan string)
    ch2 := make(chan string)

    go server1(ch1)
    go server2(ch2)

    for {
        select {
            case msg1 := <- ch1:
            	fmt.Println("Case 1:", msg1)
            case msg2 := <- ch1:
            	fmt.Println("Case 2:", msg2)
            case msg3 := <- ch2:
            	fmt.Println("Case 3:", msg3)
            case msg4 := <- ch2:
            	fmt.Println("Case 4:", msg4)
            default:
            	fmt.Println("No message")
        }
        }
    }
    
}
