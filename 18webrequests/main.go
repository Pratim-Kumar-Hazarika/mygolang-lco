package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"
 
func main(){

	fmt.Println("Welcome to web requests")
	response,err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type %T\n",response)
	fmt.Println("Status code ", response.TLS.ServerName)
	defer response.Body.Close()

	// databytes,err := io.ReadAll(response.Body)
	// checkNilErr(err)
	// content := databytes
	// fmt.Println(content)


}
func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}