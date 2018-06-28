package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Channel Examples:")
	funcname := flag.String("f", "Sumcubesquares", "funcname")
	flag.Parse()
	fmt.Println("funcname:" + *funcname)
	if *funcname == "Sumcubesquares" {
		Sumcubesquares(589)

	}

}

func Sumcubesquares(val int) {
	squarech := make(chan int)
	go Sumsquares(val, squarech)
	cubech := make(chan int)
	go Sumcubes(val, cubech)
	x, y := <-squarech, <-cubech
	fmt.Println("Sum is:", x+y)
}

func Sumsquares(val int, squarech chan int) {
	fmt.Println("Inside Sumsquares")
	digitch := make(chan int)
	sum := 0
	go Getdigits(val, digitch)

	for v := range digitch {
		fmt.Println("Inside for loop of Sumsquares")
		sum += v * v
	}
	squarech <- sum
	//<-digitch
}

func Sumcubes(val int, cubech chan int) {
	digitch := make(chan int)
	sum := 0
	go Getdigits(val, digitch)
	for digit := range digitch {
		sum += digit * digit * digit
	}
	cubech <- sum

}

func Getdigits(val int, digitch chan int) {
	fmt.Println("Inside Getdigits")
	fmt.Println("Value is:", val)
	for val != 0 {
		fmt.Println("Inside for loop of Getdigits")
		digitch <- val % 10
		val /= 10
	}
	close(digitch)
}
