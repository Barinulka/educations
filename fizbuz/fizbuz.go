package fizbuz

import "fmt"

func FizBuz(number int) string {

	if number%3 == 0 && number%5 == 0 {
		return fmt.Sprintf("%s", "FizzBuzz")
	} else if number%3 == 0 {
		return fmt.Sprintf("%s", "Fizz")
	} else if number%5 == 0 {
		return fmt.Sprintf("%s", "Buzz")
	} else {
		return fmt.Sprintf("%d", number)
	}
}

func HandleFizBuz() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizBuz(i))
	}
}
