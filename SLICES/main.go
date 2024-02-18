package main

import (
	"fmt"
	"sort"
)

func main(){

	// fmt.Println("==== Slices in go Lang ====")


	// var Fruit_list = []string{"one", "two", "three"}
	// fmt.Printf("type of fruitlist is %T ", Fruit_list)


	// Fruit_list = append(Fruit_list, "four", "five")
	// fmt.Println(Fruit_list)


	// Fruit_list = append(Fruit_list[1:3])
	// fmt.Println(Fruit_list)

	// Highscore := make([]int, 4)
	// Highscore[0] = 111
	// Highscore[1] = 222
	// Highscore[2] = 333
	// Highscore[3] = 444
	// // Highscore[4]= 555

	// Highscore = append(Highscore, 999)

	// sort.Ints(Highscore)
	// fmt.Println(Highscore)
	// fmt.Println(sort.IntsAreSorted(Highscore))


	// fmt.Println(Highscore)



	// how to remove slice based on index 

	var courses = []string{"Reactjs", "python", "c", "cpp", "swift", "ruby" }
	fmt.Println(courses)
	var index int = 2;
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}