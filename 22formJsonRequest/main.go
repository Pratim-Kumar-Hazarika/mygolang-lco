package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main()  {
	fmt.Println("Welcome to form request in golang")
	PostFormRequest()
	
}

func PostFormRequest (){
	const myurl ="http://localhost:6969/users/form"
	
	///form data

	formData := url.Values{}
	formData.Add("firstname","Pratim")
	formData.Add("lastname","Hazarika")
	formData.Add("email","hpro373@gmail.com")

	response,err := http.PostForm(myurl,formData)
	checkNilErr(err)

	defer response.Body.Close()

	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}
