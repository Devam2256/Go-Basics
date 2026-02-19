package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating maps sing literal syntax (1)
	var m map[string]int
	m = map[string]int{}
	fmt.Printf("Type of Map: %T\n", m)

	// creating maps using literal syntax (2)
	m2 := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println(m2)

	// creating maps using make function
	m3 := make(map[string]int)
	m3["orange"] = 4
	m3["grape"] = 2
	fmt.Println(m3)

	// accessing map values
	fmt.Println("Number of apples:", m2["apple"])

	// checking if a key exists (comma-ok idiom)
	v, ok := m2["banana"] // here we get two values: the value associated with the key "banana" (value) and a boolean indicating if the key exists (ok), name can be anything
	if ok {
		fmt.Println("Number of bananas:", v)
	} else {
		fmt.Println("Banana key does not exist")
	}

	delete(m2, "apple") // deleting a key-value pair from the map
	fmt.Println(m2)

	clear(m3) // clearing all key-value pairs from the map
	fmt.Println(m3)

	// cloning a map creates a new map in memory with the same key-value pairs as the original map .       
	m4 := maps.Clone(m2)
	fmt.Println(m4)
	m4["banana"] = 10
	fmt.Println("Original map:", m2)
	fmt.Println("Cloned map:", m4)

	// NOTE: panic error for maps  in  go  .       
	// var m5 map[string]int
	// m5["kiwi"] = 1 // this will cause a panic error because m5 is nil and we cannot assign a value to a nil map

	// to avoid panic error we need to initialize the map before using it
	// therefore , use either make() or literal syntax to create a map before using it

	// range over  a map
	var users = map[string]int{}
	users = map[string]int{
		"alice": 30,
		"bob":   25,
		"charlie": 35,
	}
	for k, v := range users {
		fmt.Printf("User: %s, Age: %d\n", k, v)
	}
}