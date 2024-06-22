package main

import "fmt"

func main() {
	billAmount := 0
	tip := 0.0
	tipRate := 0
	total := 0.0

	fmt.Printf("What is bill amount?")
	if _, err := fmt.Scan(&billAmount); err != nil {
		panic(err)
	}

	fmt.Printf("What is the tip rate?")
	if _, err := fmt.Scan(&tipRate); err != nil {
		panic(err)
	}

	tip = float64(billAmount) * (float64(tipRate) / 100.0)
	total = float64(billAmount) + tip

	fmt.Printf("your tip is $ %.1f\n", tip)
	fmt.Printf("your total price is $ %.1f", total)
}
