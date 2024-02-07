package stringgo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLongString(prompt string) string {
	fmt.Print(prompt)

	// fmt.Scanln(&input)
	// //We get unexpected behaviour when we enter "Learn Go" as input for the Note title prompt
	// //This is because all the scan functions are designed to accept signle word input and beyond that(space seperated words) we might not get the desired functionality
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input := strings.TrimSpace(text) //Removes leading and tralinign whitespaces
	// //OR
	// input = strings.TrimSuffix(text, "\n")
	// input = strings.TrimSuffix(input, "\r") //For windows OS
	return input
}
