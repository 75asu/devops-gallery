package main

import (
	"fmt"
	"unsafe"

	geo "go-devops/geometry"

	"rsc.io/quote"
)

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	x := 10
	name := "Asutosh Panda"
	isWorking := false

	fmt.Println("\nHello World!")
	fmt.Println(quote.Go())
	fmt.Println(x, name, isWorking)
	fmt.Printf("\nType of name %T and size is %d", name, unsafe.Sizeof(name))

	a, p := rectProps(2, 4)
	fmt.Printf("\nArea is %f and Perimeter is %f", a, p)

	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
	fmt.Println(daysOfTheMonth)

	area := geo.Area(1, 2)
	diag := geo.Diagonal(1, 2)
	fmt.Println(area, diag)
}
