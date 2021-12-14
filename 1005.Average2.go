package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)
	fmt.Scanf("%f", &n3)
	fmt.Printf("MEDIA = %.1f\n", (n1*2+n2*3+n3*5)/10)
}
