package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"os"
	//"github.com/stianeikeland/go-rpio"
	"os/signal"
	"syscall"
)

func main() {

	log.Print("aaaa")
	c := make(chan os.Signal, 1)
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	soundPin := rpio.Pin(4)
	soundPin.Input() // Input mode
	go readData(soundPin)
	defer rpio.Close()

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case s := <-c:
			log.Print(s)
			os.Exit(1)

		}
	}

}

func readData(soundPin rpio.Pin) {

	for {
		res := soundPin.Read()
		log.Print(string(res))
	}
}
