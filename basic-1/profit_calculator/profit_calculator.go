package main

import (
	"fmt"
	// "math"
)

type ProfitCalculator struct {
	revenue,expense,taxRate float64
}
func (p ProfitCalculator) ebt() float64 {
	return p.revenue - p.expense
}
func (p ProfitCalculator) eat() float64 {
	return p.ebt() - (p.ebt() * p.taxRate)
}
func main() {
	var amounts ProfitCalculator
	fmt.Println("Profit Calculator")
	fmt.Println("Enter Revenue: ")
	fmt.Scanln(&amounts.revenue)
	fmt.Println("Enter Expense: ")
	fmt.Scanln(&amounts.expense)
	fmt.Println("Enter Tax Rate: ")
	fmt.Scanln(&amounts.taxRate)

	fmt.Println("Earnings Before Tax: ", amounts.ebt())
	fmt.Println("Earnings After Tax: ", amounts.eat())

	

}