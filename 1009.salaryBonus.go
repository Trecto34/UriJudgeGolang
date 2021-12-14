package main

import "fmt"

func main() {
	var name string
	var fixSalary, totalSales, bonus float64
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &fixSalary)
	fmt.Scanf("%f", &totalSales)
	bonus = totalSales * 0.15
	fmt.Printf("TOTAL = R$ %.2f\n", fixSalary+bonus)
}
