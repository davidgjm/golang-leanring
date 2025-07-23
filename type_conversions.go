package main

import (
	"fmt"
)

func main() {
	// Integer to float conversion
	var intNum int = 10
	var floatNum float64 = float64(intNum)
	fmt.Println("Integer to Float:", intNum, "=>", floatNum)

	// Float to integer conversion
	var floatVal float64 = 10.5
	var intVal int = int(floatVal)
	fmt.Println("Float to Integer:", floatVal, "=>", intVal)

	// Demonstrating precision loss
	var preciseVal float64 = 123456789.123456789
	var impreciseVal int = int(preciseVal)
	fmt.Println("Precision Loss:", preciseVal, "=>", impreciseVal)
}
