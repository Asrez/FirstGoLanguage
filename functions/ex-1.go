package main

import "sort"

import "fmt"
func main(){
	numbers := []int{23,23,23,23,5,6,6,6}

	sort.Slice(numbers, func(i , j int ) bool {
		return numbers[i] < numbers[j]
	})

	fmt.Printf("%v\n" , numbers)
}