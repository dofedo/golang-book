package main

import ( 
	"fmt"
	"time"
)

func main() {
	// data := 0
	var data int
	// data++
	//  1. retrive
	//  2. increment
	//  3. store

	go func() { data++ }()
	// time.Sleep(1*time.Second) // This is bad!
	fmt.Println(data)
	data++
	go func() { data++ }()
	// time.Sleep(1*time.Second) // This is bad!
	fmt.Println(data)

	go func() { data++ }()
	time.Sleep(1*time.Second) // This is bad!
	fmt.Println(data)

	go func() { data++ }()
	// time.Sleep(1*time.Second) // This is bad!
	fmt.Println(data)
	

	// if data == 0 {
	// 	fmt.Println("is ", data)
	// } 
	
}