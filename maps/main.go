package main

import "fmt"

func main(){
	fmt.Println("Here  we learn Map")

	languages := make(map[string]string)

	languages["Js"] = "Javascript"
	languages["Rs"] = "rust"
	languages["Ts"] = "TypeScript"


	fmt.Println("list of all languages", languages)
	fmt.Println("Js shorts for ", languages["Js"])
	
	
	delete(languages, "Ts")
	fmt.Println("list of all languages", languages)



	// loops are instresting in go lang

	for key, value := range languages{
		fmt.Printf("for key %v , value is  %v \n", key, value)
	}



}