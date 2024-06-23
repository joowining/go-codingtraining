package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputBillAmount string
	var billAmount int64
	var tip float64
	var inputTipRate string
	var tipRate int64
	var total float64
	var err error

mainLoop:
	for {
		fmt.Printf("What is bill amount?")
		if _, err := fmt.Scan(&inputBillAmount); err != nil {
			fmt.Printf("Invalid input try")
			continue mainLoop
		}

		if billAmount, err = strconv.ParseInt(inputBillAmount, 10, 32); err != nil {
			fmt.Println("please input number")
			continue mainLoop
		}

		fmt.Printf("What is the tip rate?")
		if _, err := fmt.Scan(&inputTipRate); err != nil {
			fmt.Printf("Invalid input try")
			continue mainLoop
		}

		if tipRate, err = strconv.ParseInt(inputTipRate, 10, 32); err != nil {
			fmt.Println("please input number")
			continue mainLoop
		}

		tip = float64(billAmount) * (float64(tipRate) / 100.0)
		total = float64(billAmount) + tip
		break
	}

	fmt.Printf("your tip is $ %.1f\n", tip)
	fmt.Printf("your total price is $ %.1f", total)
}
