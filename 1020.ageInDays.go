package main

import "fmt"

func main() {
	var days, years, months int
	fmt.Scanf("%d", &days)
	for days > 29 {
		if days >= 365 {
			years += 1
			days -= 365
		} else if days >= 30 {
			months += 1
			days -= 30
		}
	}
	fmt.Printf("%d ano(s)\n", years)
	fmt.Printf("%d mes(es)\n", months)
	fmt.Printf("%d dia(s)\n", days)
}
