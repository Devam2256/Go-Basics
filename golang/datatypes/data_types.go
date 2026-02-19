package datatypes

import (
	"fmt"
)

func DataTypes() {

	fmt.Println("Hello, World!")
	fmt.Println("string" + " concatenation")
	fmt.Println("5 + 3 =", 5+3)
	fmt.Println("7.0 / 2.0 =", 7.0/2.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(true && (true || false))
	fmt.Println(!true)
}