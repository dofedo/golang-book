// Every Go program must start with a package declaration.
package main

import "fmt"

// out of scope variable declaration
// have to specify type
var out_of_scope_var string = "Hello, world!"

// Struct and declare type
type person struct {
	score int
	age   int
}

// return closure
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func countFrom0to10() {
	for counter := 0; counter <= 10; counter++ {
		if counter%2 == 0 {
			fmt.Println(counter, "even")
		} else {
			fmt.Println(counter, "odd")
		}
	}
}

// Return multiple values
func returnSomeValue() (int, int) {
	return 5, 6
}

// Get some values
// Closure
func getSomeValuePrint(score float32, age ...uint8) {
	for _, value := range age {
		score += float32(value)
	}

	showResult := func(inp float32) {
		fmt.Println(inp)
	}

	showResult(score)

}

// Map objects
func mapNames2Data() {
	m1 := make(map[string]person)
	m1["key1"] = person{score: 12, age: 34}
	m1["key2"] = person{score: 13, age: 14}
	fmt.Println(m1, m1["key1"], m1["key2"])

	for key, data := range m1 {
		fmt.Println(key, data)
	}

}

// Array example
func slicesOfThings() {

	// array declaration
	var arr1 [7]uint8
	arr1[0] = 12
	fmt.Println(arr1)

	var total uint16

	for index, value := range arr1 {
		if index+1 != len(arr1) {
			arr1[index+1] = arr1[index]
		}
		fmt.Println("arr", arr1[index])
		fmt.Println(value)
		total += uint16(value)
	}
	fmt.Println(arr1)

	fmt.Println(total)

	x := [4]float64{
		98,
		93,
		77,
		82,
		// 83,
	}
	fmt.Println(x)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice3, slice4)
	fmt.Println(slice3, slice4)

}

// get user inputs
func getNumberMultiplyBy2() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

// print strings
func printSomething() {
	// first statement
	fmt.Println("Hello, World!")

	// Types & Variables
	var s1 string = "Hello, world!"
	fmt.Println(s1)

	var s2 string
	s2 = "Hello,"
	var s3 string = "world!"
	fmt.Println(s2 + " " + s3)

	// type not-specified
	s4 := "Hello, world!"
	fmt.Println(s4)

	fmt.Println(out_of_scope_var)

}

func main() {

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	switch input {
	case 0:
		printSomething()
	case 1:
		getNumberMultiplyBy2()
	case 2:
		countFrom0to10()
	case 3:
		slicesOfThings()
	case 4:
		mapNames2Data()
	case 5:
		v1, v2 := returnSomeValue()
		fmt.Println(v1, v2)
	case 6:
		getSomeValuePrint(12.2, 12, 3, 5, 6)
		getSomeValuePrint(12.4)
	case 7:
		nextEven := makeEvenGenerator()
		fmt.Println(nextEven())
		fmt.Println(nextEven())
		fmt.Println(nextEven())
		fmt.Println(nextEven())
	case 8:
		
	default:
		break
	}

}
