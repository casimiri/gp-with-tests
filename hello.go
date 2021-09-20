package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishPrefixHello = "Hello, "
const spanishPrefixHello = "Hola, "
const frenchPrefixHello = "Bonjour, "

func Hello(name string, langage string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(langage) + name
}

func greetingPrefix(langage string) (prefix string) {

	prefix = englishPrefixHello

	switch langage {
	case spanish:
		prefix = spanishPrefixHello
	case french:
		prefix = frenchPrefixHello
	}

	return
}

func main() {
	fmt.Println(Hello("world", "French"))
}
