package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println(createErrorMethod2(0))

	output , err := createErrorMethod1(0)

	fmt.Println(output, "%s\n", err)
}


func createErrorMethod1(number int) (int , error) {
	if number == 0 {
		return 0 , errors.New("number is not valid")
	}
	return number, nil	
}

func createErrorMethod2(number int) (int , error) {
	if number == 0 {
		return 0 , fmt.Errorf("number is not valid %d", number)
	}
	return number, nil	
}

