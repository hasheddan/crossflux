package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Todo struct {
	Title    string `json:"title"`
	Assignee string `json:"assignee"`
	Done     bool   `json:"done"`
}

func handleRequests() {
	log.Println("Starting Todos server...")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/create", create).Methods("POST")
	myRouter.HandleFunc("/list", list)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/app-1?charset=utf8&parseTime=True", user, password, host))
	defer db.Close()

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	db.AutoMigrate(&Todo{})
	handleRequests()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Todos app!")
	fmt.Println("Request Received: /")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy!")
}

func create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo Todo
	json.Unmarshal(reqBody, &todo)
	db.Create(&todo)
	fmt.Println("Request Received: /create")
	json.NewEncoder(w).Encode(todo)
}

func list(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
	db.Find(&todos)
	fmt.Println("Request Received: /list")
	json.NewEncoder(w).Encode(todos)
}
