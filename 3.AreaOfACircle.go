package main

import (
	"fmt"
)

func main() {
	var r float64
	fmt.Scanln(&r)
	fmt.Printf("A=%.4f\n", r*r*3.14159)
}
