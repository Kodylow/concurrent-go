package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)
		for {
			// if no clients, barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("%s sees no clients, so takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes up barber %s.", client, barber)
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop closed, send barber home and close goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s finishes cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan(fmt.Sprintf("%s leaves the barbershop", barber))
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing the barbershop for day")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)

	color.Green("Closing the barbershop for day done")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s arrives at the barbershop***", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default:
			color.Red("Waiting room full, %s leaves", client)
		}
	} else {
		color.Red("The shop is closed, so %s leaves.", client)
	}
}
