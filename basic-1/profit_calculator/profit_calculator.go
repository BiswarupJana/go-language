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
    return p.ebt() * (1 - p.taxRate/100)
}

//ratio function
func (p ProfitCalculator) ratio() float64 {
	return p.ebt() / p.eat()
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

	// fmt.Println("Earnings Before Tax: ", amounts.ebt())
	// fmt.Println("Earnings After Tax: ", amounts.eat())
	// fmt.Println("Profit Ratio: ", amounts.ratio())

	// formatedFV := fmt.Sprintf("Earnings Before Tax: %.2f\n", amounts.ebt())
	// formatedRFV := fmt.Sprintf("Earnings After Tax: %.2f\n", amounts.eat())
	// formatedPR := fmt.Sprintf("Profit Ratio: %.2f\n", amounts.ratio())
	// fmt.Println(formatedFV, formatedRFV, formatedPR)

	fmt.Printf(`Earnings Before Tax: %.2f
	Earnings After Tax: %.2f
	Profit Ratio: %.2f`, amounts.ebt(), amounts.eat(), amounts.ratio())

}