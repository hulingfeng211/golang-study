package main

import (
	"bufio"
	"encoding/base64"
	"io"

	"github.com/jacobsa/go-serial/serial"
	"log"
)

func main()  {

	qrcodeChan :=make(chan string,1)

	//exitChan:=make(chan int)
    optionUsb:=serial.OpenOptions{
		PortName: "/dev/ttyUSB0",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		MinimumReadSize: 4,

	}
	serialPort,err:=serial.Open(optionUsb)
	if err!=nil {
		log.Fatalf("err:%s",err)
	}
	defer serialPort.Close()
	log.Print("start read data")

	for {
		getQrcodeData(serialPort,qrcodeChan)
	}



	//for {
	//	select {
	//	case data:= <- qrcodeChan:
	//		log.Print(data)
	//		qrcodeContent, err := base64.StdEncoding.DecodeString(data)
	//		if err!=nil {
	//			log.Print(err)
	//		} else {
	//			log.Print(string(qrcodeContent))
	//		}
	//	case <- exitChan:
	//
	//		break
	//
	//	}
	//}


}

//读取串口扫描数据
func getQrcodeData( serialPort io.ReadWriteCloser,c chan string)  {
	reader := bufio.NewReader(serialPort)
	reply, err := reader.ReadBytes('\x09')
	if err != nil {
		panic(err)
	}
	data:= string(reply)
	qrcodeContent, err := base64.StdEncoding.DecodeString(data)
	log.Print(string(qrcodeContent))
	//return string(reply)
}
