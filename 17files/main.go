package main

import (
	"fmt"
	"io"
	"os"
)


func main()  {
	fmt.Println("Welcome to Files in golang")

	content := "Hello welcomeP"

	file,err := os.Create("./myFile.text")

     checkNilErr(err)
	 length,err := io.WriteString(file,content)

	 checkNilErr(err)
	 fmt.Println("length is :",length)

	 defer file.Close()

	 readFile("./myFile.text")
}

func readFile(filename string){
	databyte,err := os.ReadFile(filename)
	checkNilErr(err)
	
	fmt.Println("Raw databytes ",string(databyte))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}