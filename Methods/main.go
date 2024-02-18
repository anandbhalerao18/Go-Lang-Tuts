package main

import "fmt"

func main(){
	fmt.Println("Welcome to methods")

	anand := User{
		Name: "anand",
		Email: "bhalerao1895@gmail.com",
		Age: 18,
		status: true,
	}
	fmt.Printf("The details of user are = %v", anand)

	anand.GetStatus()
	anand.NewMail()



}

type User struct{
	Name string
	Email string
	Age int
	Status bool
	
}

func (u User) GetStatus(){
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@g.com"
	fmt.Println("Email of this user is = ", u.Email)
}