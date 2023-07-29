package main

import (
	"fmt"
	"time"
)

func numbers(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Written on channel %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}

	close(n)
}

func main() {
	cn := make(chan int, 3)

	go numbers(cn)

	for v := range cn {
		fmt.Printf("Read from channel: %d\n", v)
		time.Sleep(time.Millisecond * 300)
	}

	<-cn

	fmt.Println("End")
}
