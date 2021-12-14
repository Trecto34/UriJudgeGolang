package main

import (
	"fmt"
	"math"
)

func main() {
	var rad int
	var vol float64
	fmt.Scanf("%d", &rad)
	vol = (4.0 / 3.0 * 3.14159 * math.Pow(float64(rad), 3))
	fmt.Printf("VOLUME = %.3f\n", vol)
}
