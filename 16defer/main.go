package main

import "fmt"

func main()  {
	fmt.Println("Welcome to defer in Golang")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}
///world,One,Two
///print will be Hello,Two One,World  ***LIFO***
////prints 12  line -> my defer (0,1,2,3,4) [Last in first out] 43210
///Hello 43210  Two,One,World
func myDefer(){

	for i:=0;i<5;i++{
		defer fmt.Print(i)
	}
}