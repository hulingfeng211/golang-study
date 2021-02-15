package main

import (
	"bufio"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"time"
)

func main() {
	var preTime, currentTime int64
	log.Printf("%d,%d", preTime, currentTime)
	start := time.Now()
	time.Sleep(time.Duration(2) * time.Second)
	//log.Print(start.Nanosecond())
	log.Print((time.Now().UnixNano() - start.UnixNano()) / 1e6)

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
