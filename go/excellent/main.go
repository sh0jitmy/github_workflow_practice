package main

import "fmt"

func main() {
	number := 1
	fmt.Printf("%d even or Odd =%s\n",number,EvenOrOdd(number))
}

func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
