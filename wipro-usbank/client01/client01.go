// How will you return the change to a customer
package main

import (
	"fmt"
	"math"
)

func giveChange(total, paid float64) {
	if paid < total {
		fmt.Println("Not enough money given!")
		return
	}

	change := paid - total
	fmt.Printf("Total change to return: %.2f\n", change)

	denominations := []float64{20, 10, 1, 0.25, 0.10, 0.05}
	changeMap := make(map[float64]int)

	for _, denom := range denominations {
		if change >= denom {
			count := int(change / denom)
			changeMap[denom] = count
			change = math.Round((change-float64(count)*denom)*100) / 100
		}
	}

	fmt.Println("change Breakdown:")
	for _, denom := range denominations {
		if count, exists := changeMap[denom]; exists && count > 0 {
			fmt.Printf("$%.2f x %d\n", denom, count)
		}
	}
}

func main() {
	giveChange(35.70, 50)
}
