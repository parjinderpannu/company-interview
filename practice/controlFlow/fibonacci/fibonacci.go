/*
Write a fibonacci golang program
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number to generate fibonacci series: ")
	fmt.Scanln(&n)
	fmt.Println("Fibonacci Sequence using iteration")
	fibonacciIteration(n)
	fmt.Printf("Fibonacci[%d] = %d\n", n, fibRec(n))

}

// Time / space : O(n) / O(1)
func fibonacciIteration(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func fibRec(n int) int {
	if n <= 1 {
		return n
	}
	return fibRec(n-1) + fibRec(n-2)
}
