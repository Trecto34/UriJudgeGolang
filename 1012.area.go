package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	for {
		fmt.Scanf("%f", &a)
		fmt.Scanf("%f", &b)
		fmt.Scanf("%f", &c)
		break
	}
	fmt.Printf("TRIANGULO: %.3f\n", (a*c)/2)
	fmt.Printf("CIRCULO: %.3f\n", math.Pow(c, 2)*3.14159)
	fmt.Printf("TRAPEZIO: %.3f\n", ((a+b)*c)/2)
	fmt.Printf("QUADRADO: %.3f\n", math.Pow(b, 2))
	fmt.Printf("RETANGULO: %.3f\n", a*b)
}
