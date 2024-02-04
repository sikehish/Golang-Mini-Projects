//Alternate soln: (A bit advanced)

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type AppErrors = []error

func main() {
	var appErrors AppErrors
	revenue := getUserInput("Revenue", &appErrors)
	expenses := getUserInput("Expenses", &appErrors)
	taxRate := getUserInput("Tax Rate", &appErrors)

	if len(appErrors) > 0 {
		for _, err := range appErrors {
			fmt.Println(err)
		}
		panic("program error exit(1)")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(label string, appErr *AppErrors) float64 {
	var value string
	fmt.Print(label + ": ")
	fmt.Scan(&value)

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		msg := fmt.Sprintf("%v not a valid number", label)
		*appErr = append(*appErr, errors.New(msg))
	}

	if num <= 0.0 {
		msg := fmt.Sprintf("%v must be a postive value", label)
		*appErr = append(*appErr, errors.New(msg))
	}

	return num
}
