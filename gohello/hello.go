package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const italian = "Italian"

const englishHelloPrefix = "Hello "
const spanishHolaPrefix = "Hola "
const frenchBonjourPrefix = "Bonjour "
const italianCiaoPrefix = "Ciao "

/* (prefix string) corresponds to a named return value
this will create a variable named prefix in the function and assign the value "zero" to it
you can return that value by just calling return instead of return prefix

this function starts with a lowercase letter to make it not reachable from other files .go */
func greetingPrefix(lang string) (prefix string) {

	// switch statement is used instead of a long series of if statements
	switch lang {
	case french:
		prefix = frenchBonjourPrefix
	case spanish:
		prefix = spanishHolaPrefix
	case italian:
		prefix = italianCiaoPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello function MUST start with a CAPITAL LETTER if you want to export the function
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("Go", "Spanish"))
}
