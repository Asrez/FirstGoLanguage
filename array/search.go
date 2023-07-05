package main


import "fmt"
func main(){
	names := [4]string{"soheil" , "hesam" , "ali" ,"mamamd"}
	

	for _, i := range names{
		// println(i)
		if (i == "ali"){
			fmt.Println("found ali")
		}
	}


}