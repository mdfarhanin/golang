package main

import (
	"fmt"
	"strconv"
)

var result = 0

func add(a int, b int) int {
	defer printResult(result)
	return a + b
}

func subtract(a int, b int) int {
	defer printResult(result)
	return a - b
}

func multiply(a int, b int) int {
	defer printResult(result)
	return a * b
}

func divide(a int, b int) int {
	defer printResult(result)
	return a / b
}

func Calculator(a int, b int, method string) {
	switch method {
	case "add":
		result = add(a, b)
	case "subtract":
		result = subtract(a, b)
	case "multiply":
		result = multiply(a, b)
	case "divide":
		result = divide(a, b)
	}
}

func printResult(result int) {
	fmt.Println(result)
	fmt.Println("I am in deferred function - result is " + strconv.Itoa(result))
}
