package main

import "fmt"
type Employee struct{
	ID int
	Name string
	salaryInMonth int
}


func main(){

	em := Employee{ID : 1 , Name : "sohel" , salaryInMonth : 2000}

	var salary_1 = calcutSalary(em)
	var salary_2 = em.calcutSalary()

	fmt.Println(salary_1)
	fmt.Println(salary_2)
}


// this is function
func calcutSalary(em Employee) int{
	return em.salaryInMonth * 12
}

// this is method 
func (em Employee) calcutSalary() int{
	return em.salaryInMonth * 12
}
