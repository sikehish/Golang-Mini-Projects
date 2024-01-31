package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount = 1000
	// var expectedReturnRate = 5.5
	// var years = 10

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// fmt.Println(futureValue)

	//without type conversions by explicit declaration of types
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)

}
