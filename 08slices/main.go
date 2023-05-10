package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to slices class")

	var fruitList = []string {"apple","mango","pineapple"}


	fmt.Printf("The type of %T\n ", fruitList)

	fruitList = append(fruitList, "anar","avacodo")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:]) /// removes first element in array
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3]) ///Start from 1 exclude 3
	// fmt.Println("Here ",fruitList)


	// fruitList = append(fruitList[:3]) ///Start from 0 till 3-1 = 2
	// fmt.Println("Here ",fruitList)

	
	highScores := make([]int,4)
	highScores[0]=2
	highScores[1]= 1
	highScores[2]= 3
	highScores[3] = 4

	// highScores[4] = 5 //// throws error out of bond 
	fmt.Println(highScores)

	highScores = append(highScores, 5,6,7,8,9) // reallocates memeory and adds the values 
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	///all this method available in slice only
	sort.Ints(highScores)
	fmt.Println(highScores)


	///how to remove a value from slice based on index

	var courses = []string {"reactjs","golang","typescript","sql","swift"}
	///remove index 2 ( typscript)
	fmt.Println(courses)
	var indexToBeRemoved int = 2

	courses = append(courses[:indexToBeRemoved],courses[indexToBeRemoved+1:]...)
	fmt.Println(courses)
}


