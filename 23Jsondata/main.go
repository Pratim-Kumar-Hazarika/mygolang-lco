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
	// EncodeJson()
	DecodeJson()
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
func  DecodeJson()  {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs",
		"price": 299,
		"tags": ["a","b","c"]
	}
	`)
	var myData course

	checkValid := json.Valid(jsonDataFromWeb)
	if(checkValid){
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &myData)
		fmt.Printf("%#v\n",myData)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}

var  myOnlineData map[string] interface{}
json.Unmarshal(jsonDataFromWeb,&myOnlineData)
fmt.Printf("%#v\n",myOnlineData)

for k,v := range myOnlineData {
	fmt.Printf("Key is %v and value is %v and Type is: %T", k,v,v)
}
}