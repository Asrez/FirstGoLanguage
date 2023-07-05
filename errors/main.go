package main

import (
	"io/ioutil"
	"net/http"
)



func main(){
	response , err := http.Get("https://google.com")

	if err != nil{
		println("an error")
	}

	println(response.Status)

	respons , err := ioutil.ReadAll(response.Body)

	if err != nil{
		println("an error")
	}
	println(respons)

}