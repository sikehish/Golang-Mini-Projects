// Project: Same as profit calc except that we use functions

//NOTE: You can place the functions before or after the main function...There's no compulsion in placement of functions

package main

import "fmt"

func getUserInput(prompt string) float64 {
	var inp float64
	fmt.Print(prompt + ": ")
	fmt.Scan(&inp) //OR fmt.Scanf("%v", &inp)
	return inp
}

func main() {
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %v\n", ebt) //Prints value in its default format
	fmt.Printf("Profit: %g\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses               //Earnings Before Tax(EBT)
	profit := ebt - ((taxRate / 100) * ebt) //Earnings AFter Tax OR Profit
	ratio := ebt / profit                   //EBT/Profit ratio

	return ebt, profit, ratio
}
