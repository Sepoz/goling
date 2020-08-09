package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello "
const spanishHolaPrefix = "Hola "
const frenchBonjourPrefix = "Bonjour "

// Hello function MUST start with a CAPITAL LETTER if you want to export the function
func Hello(name string, lang string) string {
	if lang == french {
		return frenchBonjourPrefix + name
	}
	if lang == spanish {
		return spanishHolaPrefix + name
	}
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Go", "Spanish"))
}
