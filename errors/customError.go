package main

import "fmt"

type HttpError struct {
	StatusCode int 
	Message string
}

func main() {
	httpError := HttpError{
        StatusCode: 404,
        Message: "Not Found",
    }
    fmt.Println(httpError)
}


func (e HttpError) Error() string {
	return fmt.Sprintf("http status code: %d, message: %s", e.StatusCode, e.Message)
}