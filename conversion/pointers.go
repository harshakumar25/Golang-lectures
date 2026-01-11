package main

import "fmt"

func main() {
	fmt.Println("welcome to the worlds of pointers in golang")

	//  var ptr *int // pointer storing integer values

	//  fmt.Println("value of pointer ptr is : ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("address of myNumber is : ", ptr)
	fmt.Println("value of ptr is : ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("adress of the pointers : ", ptr)
	fmt.Println("new number of pointer is : ", *ptr)
}
