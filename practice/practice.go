package main

import "fmt"

func main() {
	prices := map[string]float64{
		"Apple":  1.25,
		"Banana": 0.75,
		"Milk":   3.50,
		"Bread":  2.00,
	}

	prices["Eggs"] = 4.00

	displayMap(prices)
	findKeyFromMap(prices, "Eggs")
	findKeyFromMap(prices, "yogurt")

}

func displayMap(prices map[string]float64) {
	fmt.Println("Item Prices:")
	for item, price := range prices {
		fmt.Printf("%s: $%.2f\n", item, price)
	}
}

func findKeyFromMap(prices map[string]float64, item2Check string) {
	price, exists := prices[item2Check]
	if exists {
		fmt.Printf("\n%s is available at %.2f\n", item2Check, price)
	} else {
		fmt.Printf("\n%s is not available\n", item2Check)
	}
}
