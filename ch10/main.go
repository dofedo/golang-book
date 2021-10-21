package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f1(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, 'a', i)
		atm := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * atm)
	}
}

// pingOnChannel is "sender"
// for specify the type of chan use this syntax
func pingOnChan(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// pongOnChannel is both sender and receiver
// bi-directional chan
func pongOnChan(c chan string) {
	for i := 0; ; i++ {
		msg := <-c
		c <- "pong " + msg
	}
}

// lookChan is "receiver"
func lookChan(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		// time.Sleep(time.Second * 1)
	}
}

// select chan syntax
func selectChan() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func main() {

	// Ex 1
	// go f1(0)
	// go f1(1)

	// Ex 2
	// for i := 0; i < 10; i++ {
	// 	go f2(i)
	// }

	// Ex 3
	// Goroutines & Channels
	// var c chan string = make(chan string)

	// go pingOnChan(c)
	// // test reading from chan
	// go pongOnChan(c)
	// go lookChan(c)

	// select functionality
	selectChan()

	// main goroutine terminates the code before other goroutines do anything
	var input string
	fmt.Scanln(&input)
}
