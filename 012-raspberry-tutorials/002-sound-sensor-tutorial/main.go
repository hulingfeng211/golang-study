package main

import (
	"bufio"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func main() {
	optionUsb := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	serialPort, err := serial.Open(optionUsb)
	if err != nil {
		log.Fatalf("err:%s", err)
	}
	defer serialPort.Close()
	log.Print("start read data")

	reader := bufio.NewReader(serialPort)

	buffer := make([]byte, 64)
	for {
		len, err := reader.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("read len:%d", len)
		log.Printf("data:%s", string(buffer))
	}
}
