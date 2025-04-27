package main

import "fmt"
import "rsc.io/quote"

const (
	spanish = "Spanish"

	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Alexey", ""))
}