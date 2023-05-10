package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

	 readFile("./myFile.text")
}

func readFile(filename string){
	databyte,err := ioutil.ReadFile(filename)
	checkNilErr(err)
	
	fmt.Println("Raw databytes ",string(databyte))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}