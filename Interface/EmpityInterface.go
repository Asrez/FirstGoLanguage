package main


import "fmt"

func main(){
	print("hello")
	print(4)
	print(true)
}


func print(input interface{}){
	fmt.Printf("%T : %v\n" , input , input)
}