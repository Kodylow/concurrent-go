// This is a simple demo of how to solve the Sleeping Barber problem.
// A concurrency problem proposed by Dijkstra to demonstrate how
// you might not need mutexes for certain concurrency problems

// We have a finite number of barbers, a finite number of seats in a
// waiting room, a fixed length of time the barbershop is open, and
// clients arriving at roughly regular intervals. When a barber has nothing
// to do, he can check the waiting room for new clients, and if one or
// more is there can give a haircut. Otherwise, the barber goes to sleep
// until a new client arrives.

// The Barbershop Rules
// 1. If there are no customers, the barber falls asleep in the chair
// 2. A customer must wake the barber if the barber is asleep
// 3. If a customer arrives while the barber is working, the customer leaves
// if all chairs are occupied but sits in an empty chair if available
// 4. When the barber finishes a haircut, he inspects the waiting room
// to see if there are any waiting customers and falls asleep if none
// 5. Barbershop can stop accepting clients at closing time, but the
// barbers cannot leave until the waiting room is empty
// After the shop is closed and there are no clients left in the waiting
// area, the barber goes home
package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// vars
var seatingCapacity = 10
var arrivalDate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}
	color.Green("Barbershop open!!")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Steve")
	shop.addBarber("John")
	shop.addBarber("James")
	shop.addBarber("Richard")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMillis := rand.Int() % (2 * arrivalDate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillis)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	<-closed
}
