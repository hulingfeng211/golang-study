package main

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

func demo2() {
	//只读通道
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received:%d\n", result)
		}
		fmt.Println("Done receiving")
	}
	results := chanOwner()
	consumer(results)
}

func demo1() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			log.Print("enter for i:=range data")
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
}

func demo3() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buffer bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buffer, "%c", b)
		}
		fmt.Println(buffer.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}
func main() {
	demo3()
	fmt.Println("Chapter 004")
}
