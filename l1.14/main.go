package main

import (
	"fmt"
)

func detectType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan bool, chan int, chan string:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)

	fmt.Println(detectType(a)) // int
	fmt.Println(detectType(b)) // string
	fmt.Println(detectType(c)) // bool
	fmt.Println(detectType(d)) // chan
}