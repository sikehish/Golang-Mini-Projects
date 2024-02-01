package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses //Earnings Before Tax(EBT)

	profit := ebt - (taxRate * ebt) //Earnings AFter Tax OR Profit

	ratio := ebt / profit //EBT/Profit ratio

	fmt.Printf("EBT: %f\n", ebt)
	fmt.Printf("Profit: %f\n", profit)
	fmt.Printf("Ratio: %f\n", ratio)

}
