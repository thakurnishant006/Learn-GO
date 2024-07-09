package main
import "fmt"
func main(){
	var arrSlice []int =[]int{13,3,4,5}
	slicNew:=[]int{1,2,3,4,5,6,7,8}
	slicNew=slicNew[5:]
	fmt.Printf("array1 %v and array2 %v",arrSlice,slicNew)
}
// Slices are similar to array but their size can change. They are used more than array.