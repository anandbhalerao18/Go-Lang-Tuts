package main

import "fmt"

func main() {
	fmt.Println("Welcome to  arrays in go lang")

	var fruitList [4]string 
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[2] = "santra"
	fruitList[3] = "Mosambi"

	fmt.Println("Fruit List is :", fruitList)
	fmt.Println("Fruit List is :", len(fruitList))




	var veg_list = [3]string{"potato", "beans","mashroom"}
	fmt.Println("veggi list is :", veg_list)

}