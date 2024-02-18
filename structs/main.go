package main

import "fmt"

func main(){
	fmt.Println("Structs in Go lang")
	// no inheritance in go lang


	anand := User{
		Name: "anand",
		Email: "anand@gmail.com",
		Status: true,
		Age: 18,
	}
	fmt.Println(anand)
	fmt.Printf("Anand details are = %v \n ", anand)
	fmt.Printf(" Name is %v Email is %v Status is %v and Age is %v  \n ", anand.Name, anand.Email, anand.Status, anand.Age)

}


type User struct{
	Name string
	Email string
	Status bool
	Age int

}