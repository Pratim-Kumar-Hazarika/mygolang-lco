package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)


func main()  {
	fmt.Println("Welcome to GET requets")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:6969"

	response,err := http.Get(myurl)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status Code ",response.StatusCode)
	fmt.Println("Content Length is :",response.ContentLength)

	var responseString strings.Builder

	content ,_ := io.ReadAll(response.Body)
	byteCount,_ :=responseString.Write(content)


	fmt.Println(string(content))
	fmt.Println(content)
    fmt.Println(byteCount)
	fmt.Println(responseString.String())

}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}
