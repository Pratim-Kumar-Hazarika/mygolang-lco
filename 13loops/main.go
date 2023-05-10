package main

import "fmt"

func main()  {
	fmt.Println("Welcome to loops in golang")

	days := []string {"Monday","Tuesday","Wednesday","Friday","Saturday"}

	fmt.Println(days)

	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index,value := range days {
	// 	fmt.Printf("index is %v and value is %v\n",index,value)
	// }


	roughValue :=1

	for roughValue <10{
		if(roughValue ==2){
			goto pizza
		}

		if(roughValue ==5){
			break
		}
		fmt.Println("Value : ",roughValue)
		roughValue++
	}

	///goto statements
	pizza:
		fmt.Println("I love pizzzaaaa")
}