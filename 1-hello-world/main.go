package main

import (
	"fmt"

	"github.com/cyrilpanicker/golang/hello-world/package1"
)

var p = 9 //package level variable

//p:=8 -> not allowed

func main() {

	package1.Method()

	var x int //assigns zero-value
	fmt.Println(x)

	//var k -> not allowed, expecting type

	var y = 4
	fmt.Println(y)

	z := 7 //shortcut for declaration and assignment
	fmt.Println(z)

	x = 0
	y = 1
	z = 2
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// p := 8 //hides package level 'p'
	fmt.Println(p)

	for index := 0; index < 10; index++ {
		if index%2 == 0 {
			fmt.Println(index)
		}
	}

}
