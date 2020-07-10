package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// Demo1 单核CPU的的goroutine演示
func Demo1() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
func Demo2() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
func printPrime(prefix string) {
	defer wg.Done()
	//next:

	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue //next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed  ", prefix)
}

var counter int

func Demo3() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()

		value++
		counter = value
	}
}

func main() {
	//Demo2()
	Demo3()
}
