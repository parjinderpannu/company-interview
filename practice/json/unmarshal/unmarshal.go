package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     *int   `json:"age,omitempty"` // Uses pointer to check nil
	Address string `json:"address"`
}

func main() {
	jsonData := `{"name":"Alice", "address":"NY"}`
	var p Person

	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("\nValue : [%v]", p)
}
