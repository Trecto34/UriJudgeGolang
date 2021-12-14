package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)
	fmt.Printf("MEDIA = %.5f\n", (n1*3.5+n2*7.5)/11)
}
