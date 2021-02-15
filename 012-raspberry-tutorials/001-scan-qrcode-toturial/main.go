package main

import (
	"bufio"
	"encoding/base64"
	"io"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"log"
)

func main() {

	//qrcodeChan :=make(chan string)

	//exitChan:=make(chan int)
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

	//qData:="hY2NvdW50IjoiMTIyNDc4NDMzNjg5OCU2MzQwMDcwNTU5OTU5NSwxMzIxMjcxOTc2MzgzMjIxNzYxLDEzQ4NjkwODAzQwOSwxMjk0MDkwMzkwidGltZSI6MTYxMjc0ODAxNTAwMHJhY2NvdW50IjoiMTMxMDc1NjUyNDc4NDMzNjg5OCwxMzExMTg2LDEzMjU2MzDEsMTMyMTI3MTU5NjUxMzQ5NzA4OSwx2MTQ4NjkwODAzMzAyNDIsMTI5NDA4OTSwxMjk0MDkwMzk4MdGltZSI6MTYxMjc0AwMH0="
	//qrcodeContent1, err := base64.StdEncoding.DecodeString(qData)
	//log.Print(err)
	//log.Printf("qrcodeContent1:%s",string(qrcodeContent1))
	for {
		getQrcodeData1(serialPort)
	}
	//go getSystemInfo()
	//go getQrcodeData(serialPort,qrcodeChan)
	//for {
	//	select {
	//	case data:= <- qrcodeChan:
	//		log.Printf("chan data:%s",data)
	//		qrcodeContent, err := base64.StdEncoding.DecodeString(data)
	//		if err!=nil {
	//			log.Print(err)
	//		} else {
	//			log.Print(string(qrcodeContent))
	//		}
	//	//default:
	//	//	log.Print("wait scan qrcode")
	//	//case <- exitChan:
	//	//
	//	//	break
	//
	//	}
	//}

}

func getSystemInfo() {

	for {
		time.Sleep(time.Duration(100) * time.Second)
	}

}

//读取串口扫描数据
func getQrcodeData(serialPort io.ReadWriteCloser, c chan string) {
	reader := bufio.NewReader(serialPort)
	log.Print("getQrcodeData")
	for {
		log.Print("ReadBytes")
		reply, err := reader.ReadBytes('\x09')
		if err != nil {
			log.Print(err)
		}
		data := string(reply)
		log.Printf("origin:%s", data)
		qrcodeContent, err := base64.StdEncoding.DecodeString(data)
		log.Print(string(qrcodeContent))
		//return string(reply)
		log.Printf("origin:%s", data)
		c <- data
	}

}

//读取串口扫描数据
func getQrcodeData1(serialPort io.ReadWriteCloser) {
	reader := bufio.NewReader(serialPort)
	log.Print("getQrcodeData")
	for {
		log.Print("ReadBytes")
		reply, err := reader.ReadBytes('\x09')
		if err != nil {
			log.Fatal(err)
		}
		data := string(reply)
		log.Printf("origin:%s", data)
		qrcodeContent, err := base64.StdEncoding.DecodeString(data)
		log.Print(string(qrcodeContent))
		//return string(reply)
		log.Printf("origin:%s", data)
		//c <- data
	}

}
