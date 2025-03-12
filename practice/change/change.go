package main

import (
	"fmt"
	"sort"
)

func main() {
	items := map[string]float64{
		"Apple":  1.25,
		"Banana": 0.75,
		"Milk":   3.50,
		"Bread":  2.00,
	}

	items["Eggs"] = 4.00

	displayMap(items)
	displayMapSorted(items)
	findKeyFromMap(items, "Eggs")
	findKeyFromMap(items, "yogurt")

	//Asking user for an item to check
	var userInput string
	fmt.Printf("\nEnter the item name: ")
	fmt.Scanln(&userInput)
	findKeyFromMap(items, userInput)

}

func displayMapSorted(items map[string]float64) {
	fmt.Printf("\nSorted Items\n")
	var keys []string
	for item := range items {
		keys = append(keys, item)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s: %0.2f\n", key, items[key])
	}

}

func displayMap(items map[string]float64) {
	fmt.Println("Item Prices:")
	for item, price := range items {
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
