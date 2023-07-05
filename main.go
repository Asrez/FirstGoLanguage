package main


import "fmt"
import "strings"

func main(){
	// i , j := 3 ,4

	// var poi *int = &i
	
	// var poj *int = &j

	// fmt.Println("hello go ..")
	// fmt.Println("Adddres of i :" , poi)
	// fmt.Println("Adddres of i :" , poj)


	// fmt.Printf("hello world 😄 😄")
	
	// myrune := rune("hello world 😄 😄")

	// fmt.Println(myrune)


	// variable

	// var i int = 20
	// var s string = "soheil"

	// string method 

	full_name := "soheil khaledabadi"
	age := 18

	fmt.Printf("my name is %s my age is %d my age type is %T my binary age %b" , full_name , age , age , age)

	mystring := "this is text for golang touraine"
	

	fmt.Println(strings.Contains(mystring , "go")) // search in text
	fmt.Println(strings.ContainsAny(mystring , "go1")) // search in text
	fmt.Println(strings.Count(mystring , "go")) // get count of word in string 
	fmt.Println(strings.Cut(mystring , "go")) // cut word go from this string 

	//  split 

	mystringArray := strings.Split(mystring," ")

	fmt.Println(len(mystringArray))

}