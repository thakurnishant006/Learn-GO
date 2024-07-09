 package main
 import "fmt"
 func main(){
	x:=10
	ptr:=&x
	fmt.Println("simple",x)
	fmt.Println("ptr",ptr)
	fmt.Println("ptrpointer",*ptr)
 }
 //pontiner points to the address . Pointer is used to store the address of one variable into another variable so that we can access that others variable value and manipulate it.