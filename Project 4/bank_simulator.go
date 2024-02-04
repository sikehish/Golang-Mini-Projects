//PROJECT 4: Bank simulator using Control structures

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//FOR LOOP SYNTAX: for i:=0, i<5, i++{}
//Infinite FOR LOOP: for{}

// NOTE:Besides the for variations introduced in the last lectures, there also is another common variation
//     for someCondition {
// 		// do something ...
// 	  }
//   someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).
//   In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.
//THe above example is Basically Go's while loop as there isnt a dedicated while keyword in go

//Writing Account Balance to file
func handleWriteBalance(accountBalance float64) {
	balanceText := fmt.Sprint(accountBalance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
	//Reference: https://www.redhat.com/sysadmin/linux-file-permissions-explained
}

func handleReadBalance() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	//Ineterstingly, golang allows redeclaration of err variable but not any other variable in the above lines of code
	//Its called comma ok syntax: https://www.golinuxcloud.com/go-comma-ok-idiom/

	return balance, err
}

func main() {

	accountBalance, err := handleReadBalance()

	if err != nil {
		// fmt.Print("Error: ")
		// fmt.Println(err)
		//OR
		fmt.Println("ERROR: ", err)
		fmt.Println("------------------")
		//Ways to exit:
		//os.Exit(0) vs return vs panic("Sorry, cant continue") --> Panic also displays the source of error(file and line no) so it looks like a proper app. crash
	}

	//OR var accountBalance=10000.0 OR var accountBalance float64=10000
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice: ", choice)

		//1. Switch case statements(You dont need break statements after each case, unlike c++)
		//Thus using break statement will break out from the switch block but  you wouldnt break free from the loop
		//However, in if block, break would break free from the loop as we dont require break to jump out of if block unlike switch
		switch choice {
		case 1:
			fmt.Println("Your account balance is ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
				// return
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			handleWriteBalance(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Error! Your withdrawal amount greater than account balance.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			handleWriteBalance(accountBalance)
		default:
			fmt.Println("Goodbye!")
			// break ->it isredundant
			fmt.Println("Thanks for chosing our bank")
			return
		}

		//2. Conditional if based approach

		// //wantsCheckBalance:= choice==1
		// //OR
		// if choice == 1 {
		// 	fmt.Println("Your account balance is ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount!")
		// 		continue
		// 		// return
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount: ", accountBalance)
		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Enter withdrawal amount: ")
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount!")
		// 		continue
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Error! Your withdrawal amount greater than account balance.")
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated! New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// 	// os.Exit(0)
		// }

	}
	fmt.Println("Thanks for chosing our bank") // will execute for if based approach but not for switch
}
