package main

import "fmt"

var privateVariable int8
var PublicVariable uint16

func init() {
	privateVariable = 1
	PublicVariable := privateVariable
	fmt.Println("Hello World")
	fmt.Println(PublicVariable)
}

func main() {
}
