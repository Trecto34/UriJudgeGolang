package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, maiorAB, maiorABC int
	for {
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		fmt.Scanf("%d", &c)
		break
	}
	maiorAB = (a + b + int(math.Abs(float64(a-b)))) / 2
	maiorABC = (maiorAB + c + int(math.Abs(float64(maiorAB-c)))) / 2

	fmt.Printf("%d eh o maior\n", maiorABC)
}
