package main

import "fmt"

func main()  {
	fmt.Println("Structs in golang")
	//no inheretence, no super or parent
	pratim := User {"pratim","pratim@go.dev",true,18}
	fmt.Println(pratim)

	fmt.Printf("The value of %+v ",pratim)

	fmt.Printf("The email is %v ",pratim.Email)

	
}

type User struct {
	Name string
	Email string
	Status bool
	Age     int
}