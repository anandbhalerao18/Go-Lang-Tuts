package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files ! ")
	content := "THis needs to go in file - learnCodeOnline.dev"

	file, error := os.Create("./mylcogofile.txt")

	if error != nil {
		panic(error)
	}

	length, error := io.WriteString(file, content)

	if error != nil {
		panic(error)
	}
	fmt.Println("Length is = ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string){
	databyte, ok := ioutil.ReadFile(filename)
	if error != nil {
		panic(error)
	}
	fmt.Println("text data inside the file is \n ", string(databyte))
}

func checknilerr(){
	if error != nil {
		panic(error)
	}
}
