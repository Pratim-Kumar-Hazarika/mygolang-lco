package main

import "fmt"

const LoginToken string= "cbsjjw29e2e201xe1jmcm0e200i2ijfeknf" //Global 

func main()  {
	var username string = "pratim"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var smallNumber uint8 = 255 //256 
	fmt.Println(smallNumber)
	fmt.Printf("Variable is of type: %T \n",smallNumber)

	var smallFloat float32 = 255.29293003030221 ////returns 5 afte decimal
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	var smallFloat64 float64 = 255.29293003030221 ////more precise
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of type: %T \n",smallFloat64)


	var anotherVariable int 
	fmt.Println(anotherVariable)///returns 0
	fmt.Printf("Variable is of type: %T \n",anotherVariable) 


	var anotherStringVariable string 
	fmt.Println(anotherStringVariable) ////empty string
	fmt.Printf("Variable is of type: %T \n",anotherStringVariable)

	//implicit type

	var nickname = "pratim"
	fmt.Println(nickname)
	// nickname= 5  ****Throws error***


	/*WALRUS OPERATOR*/

	someToken :="xhwuwosmxskksd"
	fmt.Println(someToken)



	fmt.Println(LoginToken) ////empty string
	fmt.Printf("Variable is of type: %T \n",LoginToken)
}
// someToken2 :="xhwuwosmxskksd" -> Walrus operator is allowed only inside methods.....