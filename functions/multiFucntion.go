package main 

import "fmt"

func main(){
	calcPrice , calcTax  := caclulateRoomPrice("VIP" , 3 ,4)
	println(calcPrice , calcTax)

	// sumNumber := sumNumbers(3,4,5,3,2,4,6)

	// println(sumNumber)

	PrintLogs("Ali" , "mamad" , "Soheil" , "loglo" ,34,434,false)

}


// func caclulateRoomPrice(roomType string , nights int ,persionCount int) (int , float64){
// 	var price int
// 	var tax float64

// 	switch roomType{
// 	case "standard" :
// 		price = nights * 140000 * persionCount
// 	case "VIP" :
// 		price = nights * 49999 * persionCount
// 	}
// 	tax = float64(price) * 0.8

// 	return price , tax
// }

func caclulateRoomPrice(roomType string , nights int ,persionCount int) (price int , tax float64){
	switch roomType{
	case "standard" :
		price = nights * 140000 * persionCount
	case "VIP" :
		price = nights * 49999 * persionCount
	}
	tax = float64(price) * 0.8

	return
}

func PrintLogs(logs ...interface{}){
	for i , item := range logs {
		fmt.Printf("%d : %v\n" , i , item)
	}
}

// func sumNumbers(numbers ...int){
// 	var sum int = 20

// 	for index := range numbers{
// 			numbers[index] = numbers[index] * sum
// 	}

// 	return 
// }