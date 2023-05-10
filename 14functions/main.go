package main

import "fmt"


func main()  {
	fmt.Println("Welcome to functions in golang")
	hello1()
	hello2()

	result := adder(6,9)

	fmt.Println("The result is ",result)


	proResult, theStringValue := proAdder(1,2,3,4,5)
	fmt.Println("The pro additon value =",proResult)
	fmt.Println("The pro string value =",theStringValue)
}


////function signatures when u pass values with types and returns values with types
func adder(valueOne int,valueTwo int) int{
	return valueOne + valueTwo
}
func hello1(){
	fmt.Println("Hello1 function called")
}

func hello2(){
	fmt.Println("Hello2 function called")
}


func proAdder(values ...int)(int,string){
	total:= 0;

	for _,val :=range values {
		total+= val
	}
	return total,"Hello boissss"
}