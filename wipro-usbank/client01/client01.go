// How will you return the change to a customer
package main

import (
	"fmt"
	"math"
)

var changeAvailable = map[float64]int{
	20.00: 10,
	10.00: 10,
	5.00:  10,
	1.00:  10,
	0.25:  50,
	0.10:  50,
	0.05:  50,
	0.01:  101,
}

func giveChange(total, paid float64) {
	if paid < total {
		fmt.Println("Not enough money given!")
		return
	}

	change := paid - total
	fmt.Printf("Total change to return: %.2f\n", change)

	denominations := []float64{20, 10, 5, 1, 0.25, 0.10, 0.05, 0.01}
	changeMap := make(map[float64]int)

	for _, denom := range denominations {
		if change >= denom {
			count := int(change / denom)
			changeMap[denom] = count
			change = math.Round((change-float64(count)*denom)*100) / 100 // Avoid floating point precision issues
		}
	}

	fmt.Println("change Breakdown:")
	for _, denom := range denominations {
		if count, exists := changeMap[denom]; exists && count > 0 {
			fmt.Printf("$%.2f x %d\n", denom, count)
		}
	}
}

func totalChange() float64 {
	var amount float64
	for key, value := range changeAvailable {
		amount = amount + (key * float64(value))
	}
	return amount
}

func updateChange(denom float64, count int) {
	if val, exists := changeAvailable[denom]; exists && val > 0 && val >= count {
		changeAvailable[denom] += count
	} else {
		fmt.Printf("\nDenomination $%0.2f is either not available or less than %d\n", denom, count)
	}
}

func main() {
	giveChange(35.67, 50)
	availableMoney := totalChange()
	fmt.Printf("\nTotal Available Money: $%0.2f\n\n ", availableMoney)
	updateChange(20, -10)
	availableMoney = totalChange()
	fmt.Printf("\nTotal Available Money: $%0.2f\n\n ", availableMoney)

}
