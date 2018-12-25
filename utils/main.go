package utils

import (
	"fmt"
)

func PrintzeroValueAndType(value interface{}) {
	fmt.Printf("%v - %T\n", value, value)
}
