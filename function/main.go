package main

import "fmt"

func main(){
	fmt.Println("welcome to functions in go lang")
	greeter()


	Result := adder(3,5)
	fmt.Println("Result is ", Result)

	proRes, MyMsg := proAdder(2,3,4,5)
	fmt.Println("ProResult is ", proRes)
	fmt.Println("Mymsg is ", MyMsg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo;
}


func proAdder(values...int) (int, string){
	total := 0
	
	for _, val := range values{
		total += val

	}
	return total, "Hi this is pro return func"
}


func greeter (){
	fmt.Println("Namaste from Go lang ")
}



