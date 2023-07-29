package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func characters() {
	for c := 'a'; c < 'j'; c++ {
		fmt.Printf("%c ", c)
		time.Sleep(time.Millisecond * 230)
	}
}

func main() {
	go numbers() // goroutine will run concurrently to the rest of code
	go characters()
	time.Sleep(3 * time.Second)
	fmt.Println("End")
}
