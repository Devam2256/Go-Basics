package main

import (
	"fmt"
)

func main() {
	var users = [7]string{
		"some user 1",
		"some user 2",
		"some user 3",
		"some user 4",
		"some user 5",
		"some user 6",
		"some user 7",
	}

	var pointer *[]string

	a := users[0:4]
	b := users[3:7]
	pointer = &a
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(a) 
	fmt.Println(b)
	a[0] = "changed"

	b[0] = "changed"
	fmt.Println(users)
	fmt.Println(a)
	fmt.Println(b, "\n")
	
	// undertanding how slices work in  go and how they are different from  arrays  in  go  .    
	backing_arr := [7]int{1,2,3,4,5,6,7}

	s1 := backing_arr[:0]
	fmt.Println(printSliceDetails("s1", s1))

	s2 := s1[:4]
	fmt.Println(printSliceDetails("s2", s2))

	fmt.Println(printSliceDetails("s1", s1))
	fmt.Println(printSliceDetails("backing_arr", backing_arr[:]))

	// creating empty slice with literal syntax
	empty_slice1 := []int{}
	fmt.Println(printSliceDetails("empty_slice1", empty_slice1))
	
	// creating empty slice with make function
	empty_slice2 := make([]int, 0, 10)
	fmt.Println(printSliceDetails("empty_slice2", empty_slice2))

	//creating empty slice with "var" keyword
	var empty_slice3 []int
	fmt.Println(printSliceDetails("empty_slice3", empty_slice3))

	//creating empty slice with make function
	var empty_slice = make([]int,0,10)
	fmt.Println(printSliceDetails("empty_slice", empty_slice))

	// ❌ creating slice with make function but with length greater than capacity will throw error (capactiy >= length)
	var something = make([]int, 7, 7) // var something = make([]int, 7) -> here capacity is not provided so it will be same as  length  .    
	fmt.Println(printSliceDetails("something", something))

	var something_2 = make([]int, 0, 7)
	fmt.Println(printSliceDetails("something_2", something_2))

	tictactoe()

	// append on slices
	s := []int{1, 2, 3}
	fmt.Println(printSliceDetails("s", s))
	s = append(s, 7)
	fmt.Println(printSliceDetails("s", s))
	s = append(s, 8, 9, 10, 15)
	fmt.Println(printSliceDetails("s", s))
}

func printSliceDetails(name string, slice []int) string {
	return fmt.Sprintf("%s: %v, len: %d, cap: %d", name, slice, len(slice), cap(slice))
}