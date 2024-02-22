package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int 	`json:"price"`
	Password string  `json:"-"`
	Tags [ ]string    `json:"tags,omitempty"`
}

func  main(){
	fmt.Println("welcome to json ")
	EncodeJson()
}
func EncodeJson(){

	pratimCourses := []course {
		{
			"ReactJs",299, "abc123",[]string{"a","b","c"},
		},
		{
			"Dev",299, "abc123",[]string{"a","b","c"},
		},
		{
			"BES+",299, "abc123",nil,
		},

	
	}
	finalJosn, err := json.MarshalIndent(pratimCourses,"","\t")

	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJosn)



}