package main

import "fmt"

func main(){
	numbers := [5]int{4,5,32,1,45}

	numbers_copy := numbers

	numbers[0] = 100
	println(&numbers)
	println(&numbers_copy)

	fmt.Printf("numbers : %v\n" , numbers)
	fmt.Printf("numbers : %v\n" , numbers_copy)
}