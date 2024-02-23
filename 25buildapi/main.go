package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course - file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Author *Author
}
type Author struct{
	FirstName string `json:"firstname"`
}

//fake Db
var courses []Course

//middle-ware,helper-file

func (c *Course) isEmpty() bool{
	return c.CourseName ==""
}



func main(){

}
//serve home route
func serveHome(w http.ResponseWriter ,r *http.Request){
	w.Write([]byte ("<h1>Welcome  Golang api by pratim</h1>"))
  }

  func getAllCourses(w http.ResponseWriter ,r *http.Request){
	fmt.Println(("Get all Courses"))
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
  }

  func getOneCourse(w http.ResponseWriter ,r *http.Request){
	fmt.Println(("Get One Course"))
	w.Header().Set("Content-type","application/json")
	
	//grab id from request

	params :=  mux.Vars(r)
	//loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
  }

  func createOneCourse(w http.ResponseWriter ,r *http.Request){
	fmt.Println(("Create One Course"))
	w.Header().Set("Content-type","application/json")

		if r.Body == nil {
			json.NewEncoder(w).Encode("Please send some data")
		}

		///What about - {}

		var course Course
		_ = json.NewDecoder(r.Body).Decode(&course)
		if course.isEmpty(){
			json.NewEncoder(w).Encode("No data inside json")
			return
		}

		///generate unique id, string
		// append course into courses

		source := rand.NewSource(time.Now().UnixNano())
		randGenerator := rand.New(source)
		course.CourseId = strconv.Itoa(randGenerator.Intn(100))

		 // Add the course to your courses slice
		 courses = append(courses, course)

		 // Respond with the created course or a success message
		 json.NewEncoder(w).Encode(course)
  }


  func update one course(w http.ResponseWriter ,r *http.Request){
	fmt.Println(("Create One Course"))
	w.Header().Set("Content-type","application/json")

	// first - grab id from req
	prarams := mux.Vars(r)
  }