package main

import "fmt"

func main()  {
	fmt.Println("Structs in golang")
	//no inheretence, no super or parent
	pratim := User {"pratim","pratim@go.dev",true,18}
	fmt.Println(pratim)

	fmt.Printf("The value of %+v ",pratim)

	fmt.Printf("The email is %v\n",pratim.Email)

	 pratim.GetStatus()
	pratim.ManipulateEmail()

	fmt.Println("After manipulation----")
	/////This will remain same because object passes
	///the copy not the reference so thats why we need pointers
	fmt.Printf("The email is %v\n",pratim.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age     int
}

func (u User) GetStatus(){
	fmt.Println("Is use active :", u.Status)
}

func (u User) ManipulateEmail(){
	u.Email ="pratim@google.com"
	fmt.Println("Manipulated Email :", u.Email)
}