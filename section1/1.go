package main

import "fmt"

func main() {
	fmt.Println(CompareNumbers(1, 2))

	switch x := 2; x {
	case 3:
		fmt.Println("I am 3")
	case 2:
		fmt.Println("I am 2")
	case 4:
		fmt.Println("I am 4")
	}

	// while loop in go
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func CompareNumbers(i1, i2 int) (bool, int) {
	switch {
	case i1 > i2:
		fmt.Println("first number is greater than the second number")
		return false, i1 - i2
	case i2 > i1:
		fmt.Println("second number is greater than the first number")
		return false, i2 - i1
	}

	fmt.Println("numbers are equal!")
	return true, 0
}
