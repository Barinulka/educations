package main

import (
	"education/mod/fizbuz"
	"fmt"
)

func Hello() string {
	return "Hello"
}

func PrintInfo(name string, age int, status string) string {
	return fmt.Sprintf("Имя: %s, возраст: %d, статус: '%s'", name, age, status)
}

func main() {
	fmt.Println(Hello())

	fizbuz.HandleFizBuz()
}
