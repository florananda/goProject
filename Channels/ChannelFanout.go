package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}
func Producer(ch chan int) {
	for {
		sleep()
		n := rand.Intn(300)
		fmt.Println("Producer produced:", n)
		ch <- n
	}
}

func Consumer(ch chan int, name string) {
	for n := range ch {
		fmt.Println("Consumer ", name, " consumed:", n)

	}
}

func Fanout(chA, chB, chC chan int) {
	for {
		n := <-chA
		chB <- n
		chC <- n
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	go Producer(chA)
	go Consumer(chB, "B")
	go Consumer(chC, "C")
	Fanout(chA, chB, chC)
}
