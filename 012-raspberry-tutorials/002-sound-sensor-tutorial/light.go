package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stianeikeland/go-rpio"

	//"github.com/stianeikeland/go-rpio"
	"os/signal"
	"syscall"
)

func main() {

	log.Print("aaaa")
	//var lightIsOpen = false
	c := make(chan os.Signal, 1)
	lightChan := make(chan string)

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//soundPin := rpio.Pin(4)
	//soundPin.Input() // Input mode
	//soundPin2 := rpio.Pin(5)
	//soundPin2.Input() // Input mode

	redPin := rpio.Pin(24)
	redPin.Output()
	redPin.Write(rpio.High)

	greenPin := rpio.Pin(25)
	greenPin.Output()
	greenPin.Write(rpio.High)

	//go writeData(lightPin)
	//go turnOffLight(lightPin)

	//go readData(soundPin, clapChan)
	//	go readData(soundPin2)

	defer rpio.Close()
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//reader := bufio.NewReader(os.Stdin)
	go getInput(lightChan)
	for {

		select {
		case sig := <-lightChan:
			if strings.Compare(sig, "red") == 0 {
				//go turnOffLight(lightPin)
				//lightIsOpen = !lightIsOpen
				redPin.Write(rpio.Low)
				greenPin.Write(rpio.High)

			} else if strings.Compare(sig, "green") == 0 {
				//go turnOnLight(lightPin)
				redPin.Write(rpio.High)
				greenPin.Write(rpio.Low)
			} else if strings.Compare(sig, "all") == 0 {
				//go turnOnLight(lightPin)
				redPin.Write(rpio.Low)
				greenPin.Write(rpio.Low)
			} else if strings.Compare(sig, "exit") == 0 {
				log.Panicln("will exit light app")
				redPin.Write(rpio.High)
				greenPin.Write(rpio.High)
				os.Exit(0)
			} else {
				log.Println("unkown command,close light")
				redPin.Write(rpio.High)
				greenPin.Write(rpio.High)
			}

		case s := <-c:
			log.Print(s)
			os.Exit(1)

		}
	}

}

func getInput(c chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please Enter command: \r\n red: red is open the red light \r\n green: green is open the green light \r\n all: all is open both. \r\n exit: exit is exit app \n")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		c <- strings.TrimRight(text, "\n")
		// if strings.Compare(strings.TrimRight(text, "\n"), "red") == 0 {
		// 	log.Println("is red")
		// 	c <- "red"

		// } else if strings.Compare(strings.TrimRight(text, "\n"), "green") == 0 {

		// 	log.Println("is green")
		// 	c <- "green"

		// } else if strings.Compare(strings.TrimRight(text, "\n"), "all") == 0 {

		// 	log.Println("is green")
		// 	c <- "all"

		// } else {

		// 	log.Println("is unkown")
		// 	c <- "unkown"
		// }

	}
}
