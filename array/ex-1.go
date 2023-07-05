package main

import "fmt"

func main(){
	var myArray [4]int
	var myArray2 [2]int =[2]int{3,4}
	myArray3 := [3]int{3,3,4}
	myArray4 := [...]int{4,3}

	fmt.Println(myArray)
	fmt.Println(myArray2)
	fmt.Println(myArray3)
	fmt.Println(myArray4)

	fmt.Println(len(myArray))
}