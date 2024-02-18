package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("welcome to server in go lang");

	performGetRequest()

	
}

func performGetRequest(){
	const myURL = "http://localhost:8000/get"

	response, error := http.Get(myURL)

	if error != nil{
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println("status code : ", response.StatusCode)
	fmt.Println("content length is = ", response.ContentLength)

	content, error := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}