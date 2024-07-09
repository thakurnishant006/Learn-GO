package main

import "fmt"
func main()  {
	var name string // 1 way of declaring the variable 
	name ="Hello world"
	var( // Declaring multiple variable with type inference
		fName= "Ram"; lName = "Sharma";
		age=18
	)
	var salary int = 30000 // with initializing a value
	check:= "Go Lang" // Type Inference
	fmt.Println("1st variable",name)
	fmt.Printf("Full name is %v %v and age is %v \n",fName,lName,age) // Using Format specifiers
	fmt.Println("current salary :",salary)
	fmt.Println("tutorial is :",check)
}
/* 
Practice Question 1: Basic Variable Declaration and Initialization
Question:
Write a Go program that declares the following variables using a var block and initializes them with appropriate values:

city (string) with the value "New York".
population (int) with the value 8_336_817.
area (float64) with the value 468.484.
Print these variables in the format: "The city of [city] has a population of [population] and covers an area of [area] square miles."

*/