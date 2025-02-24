package main

import "fmt"

// Notice the channel as an argument
func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // send a value back to main
}

func main() {
	c := make(chan int) // create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) //start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c //Receive a value form channel
		fmt.Println("gopher", gopherID, "finished the dish")
	}
}
