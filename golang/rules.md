# some Go rules :-

1] All .go files in the same directory must declare the same package name (except _test.go, which may use package foo_test)
   So this is illegal:

    /project
    main.go          → package main
    hello_world.go   → package hello   ❌

--> Correct mental model (very important)
    In Go: 
    1) Directories define packages
    2) Files do NOT define packages independently
    3) package X means: “This file belongs to the package represented by this directory”

--> Proper directory layout

    project/
    ├── go.mod
    ├── main.go
    └── hello/
        └── hello.go

Note: ❗ A package must not have a main() unless it is package main, it is the entry point of the program just "public static void main" in Java.

2] In Go:
    -> Identifiers starting with an uppercase letter are exported
    -> Identifiers starting with a lowercase letter are package-private

3] ❌ You cannot have two package declarations in the same file. Each .go file must contain exactly one package

4] You cannot export anything declared inside a function , Because exportability applies to package-level identifiers only.

5] Go uses lexical (block-level) scoping for all identifiers (Variables, Constants, Types (primitive aliases, arrays, structs, interfaces, etc.), Functions, Methods, Type parameters)

A block is:

-> A function body
-> An if, for, switch, or select body
-> A { ... } explicit block
-> A package block

6] Go's type system allows:

-> Anonymous struct types
-> Anonymous array types
-> Anonymous function types
-> Anonymous interface types
-> Anonymous map types
-> Anonymous slice types
-> Anonymous channel types
but it can be one time use only, cannot be used again , for that use naming them using  identifiers .    

7] range works on "Array", "Slice", "String", "Map" and "Channel" and value in range is a copy not the original value, so changing it won't change the original

8] (IMP.) Difference between arrays and slices with mutability and immutability

arrays :- 

eg. 

a := [3]int{1,2,3}
b := a
b[0] = 100

b is a full copy of a.

That means:

1) Passing arrays to functions copies the entire array.
2) Large arrays become expensive.


slices :- 

A slice is not the data.
It is a small struct internally:

```
type slice struct {
    ptr *T      // pointer to backing array
    len int     // current length
    cap int     // capacity
}
```

so for this :- 

s1 := []int{1,2,3}
s2 := s1


(remaining)


9] maps can be declared in 2 ways :- (i) using make function and (ii) using literal syntax

NOTE:- we will get panic error if we try to initialize a map using "var" keyword .

10] const enums block needs "iota" for auto incrementing feature