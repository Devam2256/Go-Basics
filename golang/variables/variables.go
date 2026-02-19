package variables

import (
	"fmt"
)

func Variables() {

	var name string = "Alice"
	var age int = 30
	var height float64 = 5.5
	var isStudent bool = false

	// idiomatic grouped form of variables
	var (
    	x int    = 5
    	y string = "go"
    	z        = 3.14
	)
	fmt.Println("x:", x, "y:", y, "z:", z)

	var marks = 100.0 // type inference to float64

	grade := "A" // short variable declaration with type inference to string
	// Note: This syntax is only available inside functions. Package-level declarations require var

	var mode_of_transport string // Variables declared without a corresponding initialization are zero-valued.
	fmt.Println("Mode of Transport:", mode_of_transport)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Marks:", marks)
	fmt.Println("Grade:", grade)

	mode_of_transport = "Car"
	fmt.Println("Mode of Transport:", mode_of_transport)


	// summary of variables :- 

	// Feature						var	 :=
	// Multiple variables			✅	✅
	// Different types	 			✅	✅
	// Requires initialization		❌	✅
	// Works outside functions		✅	❌
	// Zero values assigned			✅	❌

	// Note: := is not assignment, it is short variable declaration
	//		 It does two things at once:
	//	  1) Declares variables
	//	  2) Initializes them

	a, b := 1, 2
	// a, b := 3, 4 // will throw error because a and b are already declared in the same scope
	
	a, b = 3,4 // correct way of reassigning values .
	fmt.Println(a, b)

	// var p int
	// p, q := 10, 20 // ✅ y is new, so it works fine .

	// var x1 int
	// var x2 int
	// x1, x2 := 5, 10 // ❌ no new variables on left side of :=
}
