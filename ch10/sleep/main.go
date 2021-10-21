package main

import (
	"fmt"
	"time"
)

func resleep(n int) {
	<-time.After(time.Second * time.Duration(n))
}

func main() {
	fmt.Println("Hello")
	resleep(3)
	fmt.Println("World!")
}
