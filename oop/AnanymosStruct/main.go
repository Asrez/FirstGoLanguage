package main


import "fmt"


type persion struct {
	int
	string
}

func main(){


	var apiResponse = struct{
		ResultCode int
		ResultMessage int
		Data string
	}{
		ResultCode : 0,
		ResultMessage : 12,
		Data : "message",
	}

	fmt.Printf("%v",apiResponse)


	soheil := persion{2,"soheil khaledabadi"}

	fmt.Printf("%v",soheil)


	
}