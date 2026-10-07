package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Enter your first number for your operation")
	fmt.Scan(&a)
	fmt.Println("Enter your second number for your operation")
	fmt.Scan(&b)

	fmt.Println("Enter your operation: +, -, *, /")
	var op string
	fmt.Scan(&op)

	
	switch op {
	case "+":
		fmt.Println("The result is:", add(a, b))
	case "-":
		fmt.Println("The result is:", sub(a, b))
	case "*":
		fmt.Println("The result is:", multiply(a, b))
	case "/":
		if b == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Println("The result is:", division(a, b))
	default:
		fmt.Println("Invalid operation")
	}
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	return a / b
}