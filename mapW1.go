package main 

import "fmt"

func way1 (){
	var a = make (map[string]string)
	a["name"] = "pranto"
	a["age"] = "32"

	fmt.Println(a)
}