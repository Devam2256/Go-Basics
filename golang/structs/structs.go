package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit as y is of int type asnd 0 is the zero value of int data type in go .
	v3 = Vertex{}       // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	type User struct {
		name string
		age  int
		email string
		password string
	}

	u1 := User{"xyz", 30, "email", "password"}
	u2 := User{}
	fmt.Println(u1)
	pointer := &u1
	fmt.Println(pointer)
	pointer.name = "abc"
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(v1, v2, v3, p)

	// incorrect way of assigning struct to a  variable
	// a := type Employee struct {
	// name string
	// age  int
	// email string
	// password string
	// }

	// correct way of assigning  "empty"  struct to a  variable (1)
	a1 := struct {
		name  string
		age   int
		email string
	}{}

	// correct way of assigning  "initalized"  struct to a  variable (2)
	a2 := struct {
		name  string
		age   int
		email string
	}{name: "xyz", age: 30, email: "xyz@yahoo.com"}

	// correct way of assigning struct to a  variable (3)
	// the fields can be initalized later and no need of {} at the  end
	var a3 struct {
    	name  string
    	age   int
    	email string
	}

	fmt.Println(a1, a2, a3)
}
