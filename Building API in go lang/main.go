package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// model for courses - file
type Course struct{
	CourseId string 'json:"courseid"'
	CourseName string 'json:"coursename"'
	CoursePrice int 'json:"price"'
	Author *Author 'json:"author"'

}

type Author struct{
	Fullname string 'json:"fullname'
	Website string 'json:"website"'
}


// fake DB
var courses[]Course

// middle ware, helper = in saprate file
func (c *Course)  IsEmpty() bool{
	return c.CourseId == "" && c.CourseName == ""


}

func main(){

}

// controller's = file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by AnandBhalerao</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one Course")
	w.Header().Set("Content-Type","application/json")

	// grab id from request 
	params := mux.Vars(r)


	// loop through courses, find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
			
		}
	}
	json.NewEncoder(w).encode("no course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type","application/json")


	//what if : body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")

	}

	// what about {}

	var course Course
	_, = json.NewDecoder(r.body) Decode(&course);
	if course.IsEmpty(){
		json.NewDecoder((w).Encode("Please send some data")
		
		return 
	}

	//genrate unique method
	// append course intp courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}


func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")


	//first = grab Id from Request
	params := mux.Vars(r)

	// loop, id,remove, add with my id
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return 
		}

	}

	//
}