package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("this is go mod")

	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}


func greeter(){
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to go lang series on Youtube</h1>"))
}