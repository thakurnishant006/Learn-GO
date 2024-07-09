package main
import "fmt"
func main()  {
	type person struct{
		name string
		age int
	}
	test:=person{"Raju",45}
	fmt.Printf("The name of the Person %v and the age is the %v",test.name,test.age)
}
//struct are used to create custom types with mixed datatypes.