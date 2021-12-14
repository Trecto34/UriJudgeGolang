package main

import "fmt"

func main() {
	var number, hoursW int
	var salary, total float64
	fmt.Scanf("%d", &number)
	fmt.Scanf("%d", &hoursW)
	fmt.Scanf("%f", &salary)
	total = float64(hoursW) * salary
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", number, total)
}
