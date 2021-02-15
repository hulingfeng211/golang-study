package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"os"
	"time"

	//"github.com/stianeikeland/go-rpio"
	"os/signal"
	"syscall"
)

func main() {

	log.Print("aaaa")
	var lightIsOpen = false
	c := make(chan os.Signal, 1)
	clapChan := make(chan uint8)

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	soundPin := rpio.Pin(4)
	soundPin.Input() // Input mode
	//soundPin2 := rpio.Pin(5)
	//soundPin2.Input() // Input mode
	lightPin := rpio.Pin(24)
	lightPin.Output()
	//go writeData(lightPin)
	go turnOffLight(lightPin)

	go readData(soundPin, clapChan)
	//	go readData(soundPin2)
	defer rpio.Close()

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-clapChan:
			if lightIsOpen {
				go turnOffLight(lightPin)
				lightIsOpen = !lightIsOpen
			} else {
				go turnOnLight(lightPin)
				lightIsOpen = !lightIsOpen
			}

		case s := <-c:
			log.Print(s)
			os.Exit(1)

		}
	}

}
func turnOnLight(lightPin rpio.Pin) {
	lightPin.Write(rpio.High)
}
func turnOffLight(lightPin rpio.Pin) {
	lightPin.Write(rpio.Low)
}

func readData(soundPin rpio.Pin, c chan uint8) {
	log.Print(string(soundPin))
	for {
		res := soundPin.Read()
		if res == rpio.Low {
			log.Printf("state==%d", res)
			c <- 1
			time.Sleep(time.Duration(10) * time.Millisecond)
		}

	}
}
