package main

import "fmt"

// En go solo existe for

func main() {

	//For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println("Counter: ", counter)
		counter++
	}

	//For forever
	count := 0
	for {
		fmt.Println(count)
		count++
	}
}
