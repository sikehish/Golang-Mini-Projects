//Project 7: Building a notes application(CLI) with the help of structs

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sikehish/Golang-Mini-Projects/Project7/note"
	"github.com/sikehish/Golang-Mini-Projects/Project7/todo"
)

//NOTE ON INTERFACES: you dont have to explicitly implement an interface. It happens implictly(it is implemented if the struct has the same methods defined)

// Going by the naming convention, if an interface has a single method, then we append 'er' to the method and that is the interface name(this isnt compulsory; just a convention)
type saver interface {
	Save() error
}

// type display interface {
// 	Display() error
// }

// If we want to have both Save() and Display ina single struct, there are 2 approaches to it
// // 1.
//
//	type outputtable interface {
//		Save() error
//		Display()
//	}
//
// OR
// 2. Embedded interfaces(saver interface is embedded into outputtable)
type outputtable interface {
	saver
	Display()
}

func main() {
	// var title, content, err = getNoteData()
	//OR //  title, content, err := getNoteData()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(title, content)

	title, content := getNoteData() //OR var title,content=func() OR var title, content string = func()
	todoText := getUserInput("Todo text: ")

	//Creating a new TODO
	newTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Creating a new NOTE
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Printing and saving TODO
	err = outputData(newTodo)
	if err != nil {
		return
	}
	//OR
	// newTodo.Display()
	// err = saveData(newTodo)
	// if err != nil {
	// 	return
	// }
	//OR
	// err = newTodo.Save()
	// if err != nil {
	// 	fmt.Println("Error: Todo couldn't be saved")
	// 	return
	// }

	// fmt.Println("Todo saved successfully")

	//Printing and saving NOTE
	outputData(newNote) //NO need to store the err and handle it as the function is getting terminated anyways as this is the last line
	//OR
	// newNote.Display()
	// saveData(newNote)
	//OR
	// err = newNote.Save()
	// if err != nil {
	// 	fmt.Println("Error: Note couldn't be saved")
	// 	return
	// }
	// fmt.Println("Note saved successfully")

}

// func getNoteData() (string, string, error) {
// 	title, err := getUserInput("Note title: ")
// 	if err != nil {
// 		return "", "", err
// 	}
// 	content, err := getUserInput("Note content: ")
// 	if err != nil {
// 		return title, "", err
// 	}

// 	return title, content, nil
// }

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

// func getUserInput(prompt string) (string, error) {
// 	var input string
// 	fmt.Print(prompt)
// 	fmt.Scanln(&input)

// 	if input == "" {
// 		return "", errors.New("Invalid input")
// 	}
// 	return input, nil
// }

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)

	// fmt.Scanln(&input)
	// //We get unexpected behaviour when we enter "Learn Go" as input for the Note title prompt
	// //This is because all the scan functions are designed to accept signle word input and beyond that(space seperated words) we might not get the desired functionality

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSpace(text) //Removes leading and tralinign whitespaces
	// //OR
	// input = strings.TrimSuffix(text, "\n")
	// input = strings.TrimSuffix(input, "\r") //For windows OS

	return input
}

//--------------------------------------------------------------

// Experimenting with interfaces: saveData takes an interface as an input param
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error: Note couldn't be saved")
		return err
	}

	fmt.Println("Note saved successfully")
	return err
}

func outputData(data outputtable) error {
	data.Display()

	err := saveData(data)
	return err
}

// --------------------------------------------------------------
//EXTRA INFO ABOUT INTERFACES:

// // The Special "Any Value Allowed" Type (interface{})
// func printSomething(value interface{}) {
// 	fmt.Println(value)
// }
// //printSomething(1),printSomething(1.5),printSomething("Hi")--> theyre all valid

// ---------------------------------------------------

// // Type switches
// // value.(type) can only be used in switch case
// func printSomething(value interface{}) {
// 	switch value.(type) {
// 	case int:
// 		fmt.Println("Integer: ", value)
// 	case float64:
// 		fmt.Println("Float: ", value)
// 	case string:
// 		fmt.Println("String: ", value)
// 		// default:  //optional
// 		// 	fmt.Println(value)
// 	}
// }

// ------------------------------------------

// Extracting type info from values
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	//if ok is true, then the value is indeed int else it isnt int
	//if ok is true, intVal is basically the value, else it defaults to 0
	//intVal+1 is valid, but value+1 isnt valid as Go compiler isnt sure of its type, and hence will throw an error

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	fmt.Println("Integer: ", value)
	fmt.Println("Float: ", value)
	fmt.Println("String: ", value)
}

//Task: Learn about generics: https://www.freecodecamp.org/news/generics-in-golang/
//Generics isn t widely used, but it overcomes the problems created by interface{}

// //The below func is similar to using interface{}, and it throws an error as T could be anything.invalid operation: operator + not defined on a (variable of type T constrained by any)
// // func add[T any](a, b T) T {
// // 	return a + b
// // }

// func log[T any](a, b T) {
// 	fmt.Println(a, b)
// }

// func add[T int | string | float64](a, b T) T {
// 	return a + b
// }
