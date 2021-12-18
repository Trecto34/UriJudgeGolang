package main

import "fmt"

func main() {
	var time, speed int
	var total float64
	fmt.Scanf("%d", &time)
	fmt.Scanf("%d", &speed)
	total = float64((time * speed)) / 12
	fmt.Printf("%.3f\n", total)
}
