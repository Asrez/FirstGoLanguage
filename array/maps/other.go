package main

import "fmt"

type Persion struct{
	Name string
	Age int
}

func main(){

	persion_map := make(map[string]Persion)


	persion_map["9304932"] = Persion{Name : "soheil" , Age : 19} 
	persion_map["343343"] = Persion{Name : "soheil" , Age : 19} 
	persion_map["3443435"] = Persion{Name : "soheil" , Age : 19}

	// get len map
	fmt.Printf("%d\n" , len(persion_map))
}