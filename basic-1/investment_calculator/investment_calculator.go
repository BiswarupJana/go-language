package main

import (
	"fmt"
	"math"
)

// Combine the parameters into a struct for better organization
type Investment struct {
	Amount, ReturnRate, InflationRate float64
	Years                             int
}

// Use the struct as a receiver to calculate future value
func (i Investment) FutureValue() float64 {
	return i.Amount * math.Pow(1+i.ReturnRate/100, float64(i.Years))
}

// Use the struct as a receiver to calculate future real value
func (i Investment) FutureRealValue() float64 {
	futureValue := i.FutureValue()
	return futureValue / math.Pow(1+i.InflationRate/100, float64(i.Years))
}

func main() {
	var investment Investment

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investment.Amount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&investment.ReturnRate)

	fmt.Print("Enter the inflation rate: ")
	fmt.Scan(&investment.InflationRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&investment.Years)

	fmt.Println("Future value is", investment.FutureValue())
	fmt.Println("Future real value is", investment.FutureRealValue())
}
