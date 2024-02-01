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

	//without type conversions by explicit declaration of type
	// var investmentAmount float64 = 1000
	// var years float64 = 10
	//OR

	// var investmentAmount, years float64 = 1000, 10 //WORKS ONLY FIF TYPES ARE SAME
	//OR
	// investmentAmount, years := 1000.0, 10.0
	// expectedReturnRate := 5.5 //OR var expectedReturnRate = 5.5
	//OR
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	// OR
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	//Factoring in inflation:
	const inflationRate = 2.5 //OR const inflationRate float64=2.5
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
