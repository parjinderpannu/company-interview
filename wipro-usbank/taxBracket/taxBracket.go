/*
Golang program showcase to calculate the taxes in bracket format
*/
package main

import "fmt"

// declare a struct
type Taxbracket struct {
	UpperLimit float64
	Rate       float64
}

// Initialize a struct with tax bracket
var taxbrackets = []Taxbracket{
	{25000, 0.00},
	{50000, 0.05},
	{75000, 0.10},
	{100000, 0.15},
	{125000, 0.20},
	{150000, 0.25},
	{175000, 0.30},
	{1e9, 0.35},
}

// CalculateTaxes calculate the taxes for given income
func CalculateTaxes(income float64) float64 {
	taxes := 0.00
	previousLimit := 0.00

	for _, bracket := range taxbrackets {
		if income > bracket.UpperLimit {
			taxes += (bracket.UpperLimit - previousLimit) * bracket.Rate
		} else {
			taxes += (income - previousLimit) * bracket.Rate
			break
		}
		previousLimit = bracket.UpperLimit
	}

	return taxes
}

func main() {
	incomes := []float64{40000, 50000, 250000}
	for _, income := range incomes {
		fmt.Printf("\n[$%.2f] Income : Taxes [$%.2f]\n", income, CalculateTaxes(income))
	}

}
