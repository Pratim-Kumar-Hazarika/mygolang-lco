package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://roundhightechmigration.pratim-kumarkum.repl.co"
 
func main(){

	fmt.Println("Welcome to web requests")
	response,err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type %T\n",response)

	defer response.Body.Close()

	databytes,err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes)
	fmt.Println(content)


}
func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}