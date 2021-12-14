package main

import "fmt"

func main() {
	var p1, p2, am1, am2 int
	var pr1, pr2, total float64
	for {
		fmt.Scanf("%d", &p1)
		fmt.Scanf("%d", &am1)
		fmt.Scanf("%f", &pr1)
		fmt.Scanf("%d", &p2)
		fmt.Scanf("%d", &am2)
		fmt.Scanf("%f", &pr2)

		total = float64(am1)*pr1 + float64(am2)*pr2
		fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
		break
	}
}
