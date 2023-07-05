package main


import "fmt"


func main(){
	//  create a maps

	names := make(map[string]string)
	names_2 := map[string]string{}
	var names_3 map[string]string = map[string]string{}
	// end create 

	names["full_name"] = "soheil khaledabadi"
	names_2["full_name"] = "soheil khaledabadi"
	names_3["full_name"] = "soheil khaledabadi"


	fmt.Println(names)
	fmt.Println(names_2)
	fmt.Println(names_3)

}