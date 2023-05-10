package main

import "fmt"

func main()  {
	 fmt.Println("Welcome to maps class today!!")

	 languages := make(map[string]string)

	 languages["JS"]= "Javascript"
	 languages["TS"] = "Typescript"
	 languages["GL"] = "Golang"
	 languages["AP"]= "Apple"

	 fmt.Println("Languages  ", languages)

	 delete(languages,"TS")
	 fmt.Println("After deletion ",languages) //sorted in alphabetical

	 ///loops in golang

	 for key,value := range languages {
		fmt.Printf("For key %v value is %v\n",key,value)
	 }
}