package main

import (
	"github.com/cyrilpanicker/golang/utils"
)

type count int

func main() {

	var number int = 1
	utils.PrintzeroValueAndType(number) //1 - int

	var dayCount count = 2
	utils.PrintzeroValueAndType(dayCount) //2 - main.count

	number = int(dayCount)
	utils.PrintzeroValueAndType(number) //2 - int

	dayCount = count(number)
	utils.PrintzeroValueAndType(dayCount) //2 - main.count

}
