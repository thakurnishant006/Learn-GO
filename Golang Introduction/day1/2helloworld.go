package main // Every go file must start with the package name statement. The package name main is used here. The main function should always reside in the main package.

import "fmt" //The import statement is used to import other packages. In our case, fmt package is imported and it will be used inside the main function to print text to the standard output.
func main(){ // The func keyword marks the beginning of a function. The program execution starts from the main function.
	fmt.Println("Hello World") // The Println function of the fmt package is used to write text to the standard output. package.Function() is the syntax to call a function in a package.
}