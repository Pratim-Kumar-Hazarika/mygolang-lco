package main

import "fmt"

func main()  {
	fmt.Println("Welcome to switch case in golang")

	var number int = 100

	switch number {
	case 1:
		fmt.Println("Number 1")
	case 2 :
		fmt.Println("Number 2")
		fallthrough
	case 3 :
		fmt.Println("Number 2")
	default:
		fmt.Println("default statement")
	}
	
}