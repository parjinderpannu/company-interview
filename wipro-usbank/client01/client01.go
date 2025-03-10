// You can edit this code!
// Click here and start typing.
package main

import "fmt"

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
			count:= int(change / denom)
			changeMap[denom] = count
			change = math.Round((change-float(count)
		}
	}
}

func main() {
	giveChange(47,50)
}