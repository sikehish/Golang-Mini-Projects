// Project 10:Concurrency- Running tasks in parallel
//VVI resource: https://www.freecodecamp.org/news/concurrent-programming-in-go/

// //NOTE:
// in Go, a goroutine is a lightweight thread managed by the Go runtime. The main function in a Go program runs in the main goroutine. The term "main goroutine" refers to the initial goroutine that starts when your program begins execution. It is the primary execution context where your main function and the majority of your program's logic run.

// Here's a breakdown:

//Goroutine: A goroutine is a concurrent thread of execution. It's a way to achieve concurrency in Go. Goroutines are similar to threads, but they are managed by the Go runtime, making them more lightweight and efficient.

//Main Goroutine: The main goroutine is the initial goroutine that runs when your program starts. The main function executes in this goroutine. Any code written in the main function, by default, runs in the main goroutine.

package main

import (
	"fmt"
	"time"
)

// func greet(phrase string) {
// 	fmt.Println("Hello!", phrase)
// }

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

// func slowGreet(phrase string) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// }

// // Using channel with slowGreet
// func slowGreet(phrase string, doneChan chan bool) {
// 	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
// 	fmt.Println("Hello!", phrase)
// 	doneChan <- true //Sending data to a channel(in this case, true)
// }

// Using channel with slowGreet, and using close()
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true //Sending data to a channel(in this case, true)
	close(doneChan)
	//close() function is oly used in the fucnction which takes the longest to execute
}

func main() {
	// greet("Nice to meet you!")
	// greet("How are you?")
	// slowGreet("How ... are ... you ...?")
	// greet("I hope you're liking the course!")

	// -------------------------------------------------------------

	// //Converting the above functions to goroutines
	// go greet("Nice to meet you!")
	// go greet("How are you?")
	// go slowGreet("How ... are ... you ...?")
	// go greet("I hope you're liking the course!")

	//There's one problem in the above case..we dont see the output, and that is because the goroutines are dispatached but we dont handle them, and before it could print the results the main function's execution is completed as goroutines are handled independently.The solutions is to use CHnnels
	//Also, learn about MUtex, waitGroups and other concurrency constructs by yourself which arent covered in the Project

	// -------------------------------------------------------------

	// //Channels: Using channel with a single goroutine
	// //Channels are used for communication..data is transmitted to channel. and only when value is updated/transmitted in a channel would it continue the execution of rhe remaining part of the function
	// done := make(chan bool)
	// go slowGreet("How ... are ... you ...?", done)
	// // fmt.Println(<-done) //We can also print done
	// <-done //We can also let the data flow to void,We're waiting for the channel to emit data
	// //Only after done is updated will go continue

	// -------------------------------------------------------------

	// //Working with single channel with multiple gorutines
	// //You can use a single channel with multiple goroutines
	// done := make(chan bool)
	// go greet("Nice to meet you!", done)
	// go greet("How are you?", done)
	// go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!", done)
	// <-done
	// //The above approach doesnt work as it ends as soon as any of the function calls updates the channel(Race condition issue)

	// //Hence, we can instead wait for four seperate dones, however, internally there is one done channel, but each of the goroutine has an independent channel, and so we can wait for all the four done channels to transmit some data. Lets wait for 3 additional channels
	// <-done
	// <-done
	// <-done
	// //And now it works!
	// //Note that the results are displayed in the order in which they've finished

	// -------------------------------------------------------------

	// //The above approach isnt very scalable as we need to keep track of the number of channels and enter <- done corresponding to each goroutines. Thus, we make use of slices
	// fmt.Println("----Slices approach-----")

	// dones := make([]chan bool, 4) //Initilaised with 4 nils, so we need to manually initialize each of the channel
	// //Intialization
	// for i := range dones { //OR for i,_:=range dones, where i=index, and _ is the dummy placehodler for value
	// 	dones[i] = make(chan bool)
	// }
	// //OR
	// // dones := []chan bool{
	// // 	make(chan bool),
	// // 	make(chan bool),
	// // 	make(chan bool),
	// // 	make(chan bool),
	// // }
	// //OR
	// // dones := []chan bool{}
	// // for i := 0; i < 4; i++ {
	// // 	dones = append(dones, make(chan bool))
	// // }

	// go greet("Nice to meet you!", dones[0])
	// go greet("How are you?", dones[1])
	// go slowGreet("How ... are ... you ...?", dones[2])
	// go greet("I hope you're liking the course!", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	// -------------------------------------------------------------
	//Optimizing usage of channels using a single channel:: Efficient approach
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	// for doneChan := range done {
	// 	fmt.Println(doneChan) //You can print the bool value as doneCHan stores the bool value updated by each of the func call
	// 	//Here, as mentioned earlier, we have access to 4 seperate dones passed to 4 different func calls,a lthough we've declared only one channel. A new channel is created for each function call, and now we can loop over done, to get access to the VALUE that is emitted by each of those channels, and so we need not use <- operator, as we dont have to wait. THis go construct makes it easier for us to deal with channels
	// }
	//OR
	for range done {
		//We need not print doneChan, we just have to loop,but in this case and the above doneCHan case, we get an error at the end: fatal error: all goroutines are asleep - deadlock!
		// We get this error because go doesnt when the channel is out of values and keeps looping.To solve this problem, We can use close(chanName) function in each goroutine which closes the channel. However you cant use this with greet() function as it executes befpre slowGreet. close() method should only be used in the function that takes the longest time to execute, like in this case, only in slowGreet, else , using it in greet() would result in displaying of only one greet result, resulting in premature termination of the progrm/execution
	}

	//NOTE: we dont go with this approach if we dont know which function would take the longest to execute

}
