package main

import "fmt"

const englisHelloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", englisHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Alexey"))
}