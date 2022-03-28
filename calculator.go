package main

import "fmt"

func add(a, b int) int {
	//Insert code here
	return a + b
}

func subtract(a, b int) int {
	//Insert code here
	return a - b
}

func multiply(a, b int) int {
	//Insert code here
	return a * b
}

func divide(a, b int) int {
	//Insert code here
	//consider for b = 0
	outcome := 0

	if b != 0 {
		outcome = a / b
	}

	return outcome
}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	//Insert code here
	if process == "add" {
		outcome := add(a, b)
		fmt.Printf("%d add %d is %d\n", a, b, outcome)
	}

	if process == "sub" {
		outcome := subtract(a, b)
		fmt.Printf("%d subtract %d is %d\n", a, b, outcome)
	}

	if process == "mul" {
		outcome := multiply(a, b)
		fmt.Printf("%d multiply %d is %d\n", a, b, outcome)
	}

	if process == "div" {

		if b == 0 {
			fmt.Println("The divisor cannot be a zero.")
		} else {
			outcome := divide(a, b)
			fmt.Printf("%d divide %d is %d\n", a, b, outcome)
		}
	}

}
