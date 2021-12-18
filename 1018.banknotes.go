package main

import "fmt"

func main() {
	var notas, cem, cinq, vin, dez, cinc, dois, um int
	fmt.Scanf("%d", &notas)
	fmt.Printf("%d\n", notas)
	for notas != 0 {
		if notas >= 100 {
			cem += 1
			notas -= 100
		} else if notas >= 50 && notas < 100 {
			cinq += 1
			notas -= 50
		} else if notas >= 20 && notas < 50 {
			vin += 1
			notas -= 20
		} else if notas >= 10 && notas < 20 {
			dez += 1
			notas -= 10
		} else if notas >= 5 && notas < 10 {
			cinc += 1
			notas -= 5
		} else if notas >= 2 && notas < 5 {
			dois += 1
			notas -= 2
		} else {
			um += 1
			notas -= 1
		}
	}
	fmt.Printf("%d nota(s) de R$ 100,00\n", cem)
	fmt.Printf("%d nota(s) de R$ 50,00\n", cinq)
	fmt.Printf("%d nota(s) de R$ 20,00\n", vin)
	fmt.Printf("%d nota(s) de R$ 10,00\n", dez)
	fmt.Printf("%d nota(s) de R$ 5,00\n", cinc)
	fmt.Printf("%d nota(s) de R$ 2,00\n", dois)
	fmt.Printf("%d nota(s) de R$ 1,00\n", um)

}
