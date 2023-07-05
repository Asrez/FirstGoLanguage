package main

import "fmt"
import "strings"

func main(){
	numbers := []int{32,32,23,4,5}
	names := []string{"soheil" , "hesam" ,"ali" , "mammd"}
	changeNumbers(numbers)
	// addItem(numbers)
	fmt.Println("%v" , numbers)
	ToUpperArray(names)
}

func changeNumbers(numbers []int){
	for index , _ := range numbers{
		numbers[index] = numbers[index] * 3
	}
}

func addItem(numbers *[]int){
	*numbers = append(*numbers, 6)
}


func ToUpperArray(names []string){
	for index,_ := range names{
		names[index] = strings.ToUpper(names[index])
		
		fmt.Println(names[index])
	}
}

func appendTowArray(numbers1 []int , numbers2 []int){
	numbers1 = append(numbers1,numbers2...)
}