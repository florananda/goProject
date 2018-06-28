package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}
func Producer(ch chan int, name string) {
	for {
		sleep()
		n := rand.Intn(500)
		fmt.Println("Channel ", name, " <-", n)
		ch <- n
	}
}
func Consumer(chC chan int) {
	for n := range chC {
		fmt.Println("chC consumed:", n)
	}
}

func Fanin(chA, chB, chC chan int) {
	var n int
	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC <- n
		}
	}
}

func main() {
	fmt.Println("Step1")
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	fmt.Println("Step2")
	go Producer(chA, "A")
	fmt.Println("Step3")
	go Producer(chB, "B")
	fmt.Println("Step4")
	go Consumer(chC)
	fmt.Println("Step5")
	Fanin(chA, chB, chC)
	fmt.Println("Step6")
}
