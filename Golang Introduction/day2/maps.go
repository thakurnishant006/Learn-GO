package main

import "fmt"

func main(){
	arr:= make(map[string]int)
	arr["t1"]=1
	arr["t2"]=2
	fmt.Println(arr)
	arrZ:=map[string]int{"t3":33,"t4":22}
	arrZ["t4"]=26
	fmt.Println(arrZ)
}
//Maps are like the similar datastructure like hash table it is used to store Key value pair.