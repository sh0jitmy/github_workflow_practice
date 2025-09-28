package main

import "fmt"

func main() {
	number := 1
	fmt.Println("%d even or Odd =%s",number,EvenOrOdd(number))
}

func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
