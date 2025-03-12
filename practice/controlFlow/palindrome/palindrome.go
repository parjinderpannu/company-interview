package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var input string
	// fmt.Printf("Enter a string: ")
	// fmt.Scanln(&input)
	// fmt.Printf("is %s palindrome = %t \n\n", input, isPalindrome(input))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string to check palindrome: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("[%s] palindrome", input)

}

func isPalindrome(input string) bool {
	for i := 0; i <= len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
