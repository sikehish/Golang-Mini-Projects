// Project 10:Concurrency- Running tasks in parallel

package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
}

func main() {
	// greet("Nice to meet you!")
	// greet("How are you?")
	// slowGreet("How ... are ... you ...?")
	// greet("I hope you're liking the course!")

	//Converting the above functions to goroutines
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How ... are ... you ...?")
	go greet("I hope you're liking the course!")
}
