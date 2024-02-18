package main

import "fmt"

func main(){
	defer fmt.Println("[differ line one ] This is the defer line and it is at the top of the function. this defer fucntion is executed at the last of the function even if we place the defet fucntion at the top, middle, or at the bottom of the fucntion. this will be executed at the last of the fuction \n")
	defer fmt.Println("[differ line two ] This is the defer line and it is at the top of the function. this defer fucntion is executed at the last of the function even if we place the defet fucntion at the top, middle, or at the bottom of the fucntion. this will be executed at the last of the fuction \n")
	defer fmt.Println("[differ line three ] This is the defer line and it is at the top of the function. this defer fucntion is executed at the last of the function even if we place the defet fucntion at the top, middle, or at the bottom of the fucntion. this will be executed at the last of the fuction \n")
	fmt.Println("This is the basic print line and it is below of the defer fuction \n ")

	fmt.Println("This is the last line of the fuction which is at the bottom of the whole fuction \n ")

	myDefer()
}


func myDefer(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}