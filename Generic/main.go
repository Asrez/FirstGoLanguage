package main

import "fmt"


type Numbers interface{
	int | int64 | uint64 | uint32 | uint16 | uint | float32 | float64
}
func main(){
	x := sum(2 ,4)

	fmt.Printf("%d",x)

	f := sum(2.4 ,4.0)

	fmt.Printf("%f\n",f)

	sumwithinterface := sumWithInterface(2.4 , 4.4)
	fmt.Println("%f\n",sumwithinterface)

}

func sumInteger(a , b int) int {
	return a + b
}


func sumFloat(a , b float64) float64 {
	return a + b
}


func sum[T int | float64](a , b T) T {
	return a + b
}

func sumWithInterface[T Numbers](a , b T) T{
	return a + b
}


func callApi[Input any , Output any](heder string) Output{
	
}