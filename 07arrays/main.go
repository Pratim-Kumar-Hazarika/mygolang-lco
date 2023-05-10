package main

import "fmt"

func main()  {
	fmt.Println("Welcome to array class")

	var fruitList [4]string
	fruitList[0]="apple"
	fruitList[1]="mango"
	fruitList[3]="strawbeery"

	fmt.Println("Fruit List ",fruitList)
	fmt.Println("Fruit List length ",len(fruitList)) 
	// it will return 4 not 3 :P why ? 
	//because we allocated the length to be 4.
	// Uninitialized values are assigned with 0 values in here empty strinr  "  "

	var vegList = [3]string {"tomato","aloo","kanda"}
	fmt.Println("Fruit List ",vegList)
	fmt.Println("Fruit List length ",len(vegList)) 

	//same logic as above



	
}