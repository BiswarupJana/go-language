package main

import (
	"example/prices"
	"fmt"
)

func main() {
	pricesList := []float64{10, 20, 30, 40, 50}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64]map[string]float64)

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate, pricesList)
		job.Process()
		result[taxRate] = job.TaxIncludedPrices
	}

	fmt.Println(result)
}
