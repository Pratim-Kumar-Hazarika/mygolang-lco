package main

import "fmt"

func main()  {
	fmt.Println("Welcome to if else in golangg")

	loginCount :=11

	var result string
	
	if loginCount <10 {
		result ="Login count less then 10"
	} else if(loginCount >10){
		result = "Watch out"
	}else {
		result ="Login count ==== 10"
	}


	fmt.Println(result)

	if 4%2==0 {
		fmt.Println("Even Number")
	}else {
		fmt.Println("Odd Number")
	}

}