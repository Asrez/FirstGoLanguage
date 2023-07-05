package main 

import "fmt"



type Runner interface{
	Run()
}


type Player struct{
	Name string
	Age int
	Height int
}

func main(){
	player_1 := Player{
		Name : "soheil",
		Age : 29,
		Height : 198,
	}

	// interface 
	var runner Runner = player_1
	runner.Run()

}

func (player Player)Run(){
	fmt.Printf("player is running..%s",player.Name)
}