package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {;
	welcome :="Welcommeeee 🤘🏻"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name broski")

	//comma ok, error ok

	input,_ := reader.ReadString('\n')
	fmt.Println("Your name is =",input) ////empty string
	fmt.Printf("Variable is of type: %T \n",input)

	/*
	if you enter 9 it is a string type!!!
	*/

}