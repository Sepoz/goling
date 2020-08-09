package main

import "fmt"

const englishHelloPrefix = "Hello "

// Hello function MUST start with a CAPITAL LETTER if you want to export the function
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Go"))
}
