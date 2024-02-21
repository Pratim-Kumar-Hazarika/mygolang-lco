package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our Ftv appp!!")

	fmt.Println("Please enter a number bc")
	reader :=bufio.NewReader(os.Stdin)//os.Standard input

	input,_ := reader.ReadString('\n')

	fmt.Println("You have entered =",input)

	//increase the number to input +1

    incrementedNo ,err:= strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("Yooo your rating is increased by ,",incrementedNo+5)
	}
}