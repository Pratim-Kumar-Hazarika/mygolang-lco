package main

import (
	"fmt"
	"net/url"
)


const myurl = "https://www.pratim.com:3000/courses?name=pratim&lastname=hazarika&payment_id=8991"
func main()  {
	fmt.Println("Welcome to Url in golang")

	fmt.Println(myurl)

	//parsing

	result,_ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams:= result.Query()
	fmt.Printf("The type of query params are %T\n ",qparams)
	fmt.Println(qparams["name"])

	for _, value := range qparams{
		fmt.Println(value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "pratim.com",
		Path: "/billionaire",
		RawPath: "user=xx",
	}
	anotherUrl := partsOfUrl

	fmt.Println(anotherUrl)
}