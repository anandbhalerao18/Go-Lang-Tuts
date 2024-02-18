package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://mail.google.com/mail/u/0/#inbox"

func main(){
	fmt.Println("web request in go lang")

	response, error := http.Get(url)
	if error != nil{
		panic(error)
	}
	fmt.Printf("Response is of type %T \n ", response)

	defer response.Body.Close() // coller's responsiblity to close the conection

	databytes, error := ioutil.ReadAll(response.Body)
	if error != nil{
		panic(error)
	}
	content := string(databytes)
	fmt.Println(content)

}