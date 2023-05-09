package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Time package study in golang")

	currentTime := time.Now()

	fmt.Println(currentTime)

	/*
	01-02-2006 works in the doc
	*/
	fmt.Println(currentTime.Format("01-02-2006 15:04:15 Monday"))

	/*Creating a Date*/

	createDate := time.Date(2022,time.May,28,15,5,0,0,time.UTC)

	fmt.Println(createDate)

	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

///GOOS="windows" go build
// GOOS="linus" go build
// GOOS="darwin" go build