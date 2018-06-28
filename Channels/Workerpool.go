package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(in chan int) {
	n := 0
	for {
		//n := rand.Intn(1000)
		fmt.Println("Channel send <-", n)
		in <- n
		n++
	}
}

func Consumer(out chan int) {
	for n := range out {
		fmt.Println("<-Channel Receive:", n)
	}
}

func Worker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		out <- n
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	go Producer(in)
	for i := 0; i < 100; i++ {
		go Worker(in, out)
	}
	Consumer(out)
}
