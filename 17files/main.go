package main

import (
	"fmt"
	"io"
	"os"
)


func main()  {
	fmt.Println("Welcome to Files in golang")

	content := "Hello welcome to Namaste Golang with Pratim :P"

	file,err := os.Create("./myFile.text")

     if err != nil{
		panic(err)
	 }

	 length,err := io.WriteString(file,content)

	 if err != nil{
		panic(err)
	 }
	 fmt.Println("length is :",length)

	 defer file.Close()


}