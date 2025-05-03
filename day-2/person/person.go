package main

import "fmt"

func IsAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

func main() {
	if IsAdult(18) {
		fmt.Println("Все ок")
	} else {
		fmt.Println("Все плохо")
	}
}
