package main



func main(){
	
	for {
		println("hello wold")
		break
	}

	i:= 0

	for i < 10 {
			println("hello world" , i)
			i++
	}

	for j:= 0 ; j< 10 ; j++ {
		println("hello wold !" , j)
	}

	lst := []int{1,2,3,44,43,43,43}

	for key , value := range lst { 
		println("this is value" , key , value)
	}
}