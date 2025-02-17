package main

import (
	"fmt"
	"time"
)

func greet(phrase string, donechan chan bool) {
	fmt.Println("Hello!", phrase)
	donechan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true 
	close(doneChan)
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	for range done {} 
    fmt.Println("Main function finished.")
}
