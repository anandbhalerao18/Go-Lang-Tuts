package main

import (
	"fmt"
	"net/url"
)

const myURl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456"


func main(){
	fmt.Println("This is handling URL in go lang")
	fmt.Println(myURl)


	// parsing URL = extracting information
	result, _ := url.Parse(myURl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())


	qprams := result.Query()
	fmt.Printf("the type of querry params are : %T \n ", qprams)


	// fmt.Println(qprams(["coursename"])

	for _, val := range qprams{
		fmt.Println("param is : " val)
	}
	
	partsofurl := &url.URL{
		Scheme: "https",
		host : "lco.dev",
		Path: "/tutcss",
		RawPath: "user=anand",

	}

	anotherurl := partsofurl.String()
	fmt.println(anotherurl)
}