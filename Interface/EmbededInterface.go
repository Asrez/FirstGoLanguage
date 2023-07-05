package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
	Walk()
}


type Human interface {
	Animal
	// Eat()
	// Sleep()
	// Walk()
	Speck()
}


type Persion struct {
	Name string
	Age string
}


type cat struct{
	Name string
}

func main(){

}


func (animal Animal) Eat(){
	fmt.Println("cat is eating")
}

func (animal Animal) Sleep(){
	fmt.Println("cat is sleeping")
}


func (animal Animal) Walk(){
	fmt.Println("cat is walking")
}


func (human Human) Walk(){
	fmt.Println("human is walking")
}