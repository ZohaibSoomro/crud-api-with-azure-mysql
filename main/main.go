package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "zohaib:Soomro123@tcp(golang-crud.mysql.database.azure.com:3306)/golang-crud?charset=utf8mb4&parseTime=True"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func main() {
	db.AutoMigrate(&Student{})
	router := mux.NewRouter()
	router.HandleFunc("/students", GetStudents).Methods("GET")
	router.HandleFunc("/students/{rollNo}", GetStudentWithRollNo).Methods("GET")
	router.HandleFunc("/students/create", CreateStudent).Methods("POST")
	router.HandleFunc("/students/delete/{rollNo}", DeleteStudent).Methods("DELETE")
	print("Server starting...")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type Student struct {
	*gorm.Model `json:"-"`
	RollNo      string `json:"roll_no"`
	Name        string `json:"name"`
}

var GetStudents = func(w http.ResponseWriter, r *http.Request) {
	var students []Student
	db.Find(&students)
	fmt.Printf("\nStudents: %+v\n", &students)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

var GetStudentWithRollNo = func(w http.ResponseWriter, r *http.Request) {
	var student Student
	rollNo := mux.Vars(r)["rollNo"]
	db.Find(&student, "roll_no", rollNo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}
var CreateStudent = func(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	fmt.Println("student creation: ", student)
	if err != nil {
		http.Error(w, "Invalid payload!", http.StatusNotAcceptable)
		return
	}
	db.Create(&student)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}
var DeleteStudent = func(w http.ResponseWriter, r *http.Request) {
	var students []Student
	rollNo := mux.Vars(r)["rollNo"]
	print("rollno: ", rollNo)
	db.Where("roll_no = ?", rollNo).Delete(&students)

	db.Find(&students)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}
