package main

import (
	"encoding/json"
	"fmt"
)

type testJson struct {
	name string
	family string
}

func main() {
	json_test := testJson { name : "test", family : " testing" }
	fmt.Println(json_test)
	test , _ := json.Marshal(json_test)

	println(string(test))

}