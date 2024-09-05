package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	result := x
	for math.Abs(result*result-x) > 0 {
		result = result - (result*result-x)/(2*result)
		fmt.Println("loop: ", result)
	}
	return result
}
func main() {

	fmt.Println(Sqrt(9))
}
