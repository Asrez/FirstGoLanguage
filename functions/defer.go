package main 


import "fmt"
import "os"
import "io"
func main(){

	defer println("hello wold")

	println("god bay")

	// CopyFile()
}

func CopyFile(dis , srcName string) (written int64 ,err error){
	source , err := os.Open(srcName)

	if err != nil{
		return 

	}
	defer source.Close()

	des , err := os.Create(dis)
	

	if err != nil{
		return
	}
	io.Copy(des , source)

	defer source.Close()

	return io.Copy(des , source)
}