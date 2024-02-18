package main

import "fmt"

var jwtToken = 9000000 // this is how to decleare variable out of func
const LoginToken string = "kdhfoaehf" // when we declare variable with capital Letter at first of the variable in go,so it is public variable


func main() {
	var username string = "Anand"
	fmt.Println(username)
	fmt.Printf("variables is of type %T \n ", username)


	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("isloggedin is of type %T \n ", isloggedin)


	//default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar);
	fmt.Printf("Variabe is of type %T \n ", anotherVar);

	// implecite type 
	var website = "google.com";
	fmt.Println(website);

	// no var style
	Numberofuser := 9000000
	fmt.Println(Numberofuser)

	fmt.Println(LoginToken);
	fmt.Printf("variable is of type %T", LoginToken)

}