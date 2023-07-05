package main


import "fmt"
type Persion struct{
	firstName string
	lastName string
	age int
}

func main(){

	//  1
	p := Persion{firstName : "soheil" , lastName : "khaledabadi" , age : 20}
	fmt.Println(p)
	// 2
	
	persion_2 := new(Persion)
	persion_2.firstName = "hesam"
	persion_2.lastName = "khaledabadi"
	persion_2.age = 12
	fmt.Println(persion_2)

	// 3
	persion_3 := &Persion{firstName : "mahsa" , lastName : "shafiyi" , age : 22}

	fmt.Println(persion_3)

	// 4

	persion_4 := CreateNewPersion("ss" , "soheili" , 28)

	fmt.Println(persion_4)
}

// 4

func CreateNewPersion(fistname string , lastname string , age int) *Persion{
	if len(fistname) < 3 {
		return nil
	}
	return &Persion{firstName : fistname , lastName : lastname , age : age}
}