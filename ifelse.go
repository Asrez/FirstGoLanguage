package main 

import "fmt"

func main(){
	var salary float64

	var minSalary float64

	var tax float64 = 0

	fmt.Print("enter your salary : ")
	fmt.Scanln(&salary)

	if salary <= minSalary {
		tax = 0
	}else if salary > minSalary && salary <= minSalary * 2{
		tax = 0.5
	}else{
		tax = 0.15
	}

	fmt.Printf("your tax payment : %f\n" , tax)
	fmt.Printf("your tax is : %.2f\n" , tax * salary)

}