package main

import "fmt"

func main() {
	//Question 1
	Main1()

	//Question 2
	Calculator(10, 20, "add")
	Calculator(10, 20, "subtract")
	Calculator(10, 20, "multiply")
	Calculator(100, 20, "divide")

	//Question 3
	fmt.Println(Recursive("This is a recursion call function"))

}
