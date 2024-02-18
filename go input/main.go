package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza = \n ")

	// comma ok syntex || error ok syntax
	input, _  := reader.ReadString('\n')
	fmt.Println("thanks for rating ! \n ", input)
	fmt.Printf("Type of this rating is ,  %T", input)
}