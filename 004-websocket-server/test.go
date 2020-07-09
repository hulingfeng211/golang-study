package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("进入for循环")
		fmt.Println(len(c))
		select { //如果select case的通道已满，case语句会被阻塞，多个case语句只要有一个满足运行条件，则正常运行
		case c <- x: //通道写入后阻塞在此等待其他协程读出(不带缓冲区的chan)
			fmt.Println("进入通道写")
			x, y = y, x+y
		case single := <-quit:
			fmt.Println("exit value", single)
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("准备读通道c", time.Now())
			//todo
			fmt.Println(<-c) //通道读
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
