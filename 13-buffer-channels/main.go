package main

import (
	"fmt"
	"time"
)

func numbers(done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}

	done <- true
}

// chan<- : only write channel
func characters(done chan<- bool) {
	for c := 'a'; c < 'j'; c++ {
		fmt.Printf("%c ", c)
		time.Sleep(time.Millisecond * 230)
	}

	// channel<- value: write on channel
	done <- true
}

func main() {
	// chan (without <-): write and read channel
	cn := make(chan bool)
	cl := make(chan bool)

	go numbers(cn)
	go characters(cl)

	// <-channel: read it
	<-cn
	<-cl

	fmt.Println("End")
}
