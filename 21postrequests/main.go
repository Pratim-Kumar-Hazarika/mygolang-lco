package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Welcome to post request in golang")
	PostRequest()
	
}

func PostRequest (){
	const url ="http://localhost:6969/users"
	
	///Json payload
	requestBody := strings.NewReader(`
	 {
		"name":"pratim",
		"last_name":"billionare"
	 }
	`)
	response,err := http.Post(url,"application/json",requestBody)
	checkNilErr(err)
	defer response.Body.Close()

	content ,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}
