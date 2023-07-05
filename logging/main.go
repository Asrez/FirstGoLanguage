package main

import (
	"log"
	"os"
)

func init() {
	file , err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY , 0666)
	// file , err := os.Open("log.txt")
	if err != nil {
	log.Fatalf("Could not open log error: %v", err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	log.Println("starting...")
	sum(3,5)
}

func sum(a , b int) {
	log.Println("start of sum")
	println(a + b)
	log.Println("end of sum")
}