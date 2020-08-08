package main

import "fmt"

// Hello function MUST start with a CAPITAL LETTER if you want to export the function
func Hello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(Hello("Go"))
}
