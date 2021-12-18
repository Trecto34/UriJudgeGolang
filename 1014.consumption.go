package main

import "fmt"

func main() {
	var km int
	var fuel, total float64
	fmt.Scanf("%d", &km)
	fmt.Scanf("%f", &fuel)
	total = float64(km) / fuel
	fmt.Printf("%.3f km/l\n", total)

}
