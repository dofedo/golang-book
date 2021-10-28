package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func main() {

	a1 := Author{Name: "sam", Age: 12}

	b1 := Book{
		Title:  "b1 test",
		Author: a1,
	}

	fmt.Println(b1)

	// Unmarshal
	byte_arr, err := json.MarshalIndent(b1, "", "  ")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(byte_arr))

}
