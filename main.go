package main

import "fmt"



func main(){

	fmt.Println("Please provide a function: ")
	var userInput string
	fmt.Scanf("%s", &userInput)
	
	fmt.Println(Parser(userInput))
}
