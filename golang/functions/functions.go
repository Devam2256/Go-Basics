package main

import (
	"fmt"
)

func main() {
	printValues(5, 10)
	
	result := add(3, 4)
	fmt.Println("Result of addition:", result)
	statusCode, msg := print_status_code(200, "OK")
	fmt.Println("Status Code:", statusCode, "Message:", msg)

	sum(1, 2, 3, 4, 5)
	
	anonymousFunc()

}

// func with parameters 
func printValues(a, b int) {
	fmt.Println(a, b)
}

// function with  single return value
func add(a, b int) int {
	return a + b
}

// with multiple return values
func print_status_code(status_code int, msg string) (int, string) {
	return status_code, msg
}

// variadic functions - functions that can accept a variable number of arguments of the same type
func sum(nums ...int) {
	// Within the function, the type of nums is equivalent to []int
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

// anonymous function - a function without a name, often used as a value or passed as an argument to another function
var anonymousFunc = func() {
	fmt.Println("This is an anonymous function")
}

// higher-order function - a function that takes another function as an argument or returns a function as a result
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// closure - a function that captures and retains access to variables from its surrounding scope, even after that scope has finished executing
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}