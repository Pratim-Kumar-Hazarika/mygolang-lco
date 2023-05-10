package main

import "fmt"

func main()  {
	fmt.Println("Welcome to my pointer class")

	var ptr *int
	fmt.Println("My ptr value ",ptr) ////returns nil

	myNumber :=25

	var newPtr = &myNumber ////creates a reference points to the address of the memory

	fmt.Println("newPtr===",newPtr) /// prints address
	fmt.Println("*newPtr ",*newPtr)

	*newPtr = *newPtr + 10
	fmt.Println("After addition ",myNumber)


}