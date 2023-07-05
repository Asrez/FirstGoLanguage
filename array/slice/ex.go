package main 

import "fmt"


func main(){
	myArray := [5]int{32,32,23,4,5}

	// mySlice :=[]int{3,4}

	// mySlice_1 := make([]int , 8)

	// mySlice_2 := make([]int , 8 , 16)

	slc2 := myArray[:3]

	fmt.Println("%v\n",slc2)


}

