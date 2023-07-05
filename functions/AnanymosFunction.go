package main


import "fmt"


func main(){
	func(){
		fmt.Println("Hello world")
	} () // with this call function


	myFucntion := func(){
		fmt.Println("this is my function")
	}

	myFucntion()


	println(func(numbers ...int) int {
		var total int
		for _,item := range numbers{
			total += item
		}
		return total
	}(3,5,2,5))
}