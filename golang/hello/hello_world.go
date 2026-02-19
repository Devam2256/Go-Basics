package hello

import (
	"fmt"
)

// capitalize the first letter of the function name to make it exported and accessible from other packages otherwise it remains package-private
func SayHello() {
	// hello world
	fmt.Println("Hello, World!")

	// Note: Semicolons are not required in idiomatic Go code. The Go language specification formally uses semicolons to terminate statements, but a simple automatic semicolon insertion (ASI) rule is applied by the lexer, making explicit semicolons unnecessary in almost all practical cases.
	fmt.Println("Semicolons are not required in Go code but can be used and don't give , any errors");
}
