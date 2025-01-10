// Building api in golang
// We are making seperate folder's out there but soon we'll.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // Here we are not creating copy of Author we are referencing the real one.
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB - Slice
var courses []Course

// Middleware, helper - file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building API in golang")

	router := mux.NewRouter()

	// seeding
	courses = append(
		courses,
		Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Aditya Pawar", Website: "aditya.go.dev"}},
		Course{CourseId: "4", CourseName: "Jenkins", CoursePrice: 899, Author: &Author{Fullname: "Yeshank Pawar", Website: "yeshank.go.dev"}},
	)

	router.HandleFunc("/", ServeHome).Methods("GET")
	router.HandleFunc("/courses", GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	router.HandleFunc("/course", CreateOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
	log.Println("Server is started at 4000 port")
}

// Controllers - Controller handle the request - file

// Serve Home route
// Don't interchange the paramters position of this method.
func ServeHome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Welcome to the golang api</h1>"))
}

// Note - convention -> getCourses
func GetAllCourses(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get all courses")
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(courses) // By this method we are giving response and it takes writer and and encode the value into json and throw back the response whoever is consuming it.
}

// Note - convention -> getCourse
func GetOneCourse(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get one course")
	writer.Header().Set("Content-Type", "application/json")

	// Grab id from request.
	params := mux.Vars(request)

	// loop through the courses find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(writer).Encode(course)
			return
		}
	}
	json.NewEncoder(writer).Encode("No course found with given id")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// Checking for duplicate course id
	for _, course := range courses {
		if course.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicate course id")
			return
		}
	}

	// generate unique id, string
	// append course into courses

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	// First - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Id is not found")
	return
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one  course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, match, remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Id is not found")
	return
}
