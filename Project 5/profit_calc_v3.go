//Project 5: Enhanced profit calculator

//Goals:
//1. Validate user input: Show error messages and exit if invalid input is provided, Nonegative numbers and 0s as inputs
//2. Store calculated results into a file

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	expenses, err2 := getUserInput("Expenses: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("Tax Rate: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//Approach 2: Checking for errors at the end, instead of checking for err after each input(which ise ffective  but redundant at the same time)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Invalid inputs!")
		return
		//OR panic("Invalid inputs")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	//Stroing results in a file
	storeResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Only non-negative numbers are allowed")
	}

	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	businessText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("business.txt", []byte(businessText), 0644)
}
