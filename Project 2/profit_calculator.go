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

	profit := ebt - ((taxRate / 100) * ebt) //Earnings AFter Tax OR Profit

	ratio := ebt / profit //EBT/Profit ratio

	//NOTE: %f prints the corresponding number as a decimal floating point number (e.g. 321.65), %e prints the number in scientific notation (e.g. 3.2165e+2). Generally speaking, %g strips any trailing zeroes and trailing decimal point from these two representations and prints whichever is shortest.

	fmt.Printf("EBT: %v\n", ebt) //Prints value in its default format
	fmt.Printf("Profit: %g\n", profit)
	fmt.Printf("Ratio: %f\n", ratio)

}
