package main

import (
	"fmt"
	"sync"
	"time"
)

// Number of times to eat the meal
// const feeding = 3

var period = struct {
	think int
	eat   int
	wait int
}{
	think: 2,
	eat:   3,
	wait: 2,
}

// var sticks = map[int]string {
// 	1: "free",
// 	2: "free",
// 	3: "free",
// 	4: "free",
// 	5: "free",
// }

var wg sync.WaitGroup

func main() {

	// Set the table - Hypothetical
	fmt.Println("The table is set!")
	
	
	// The philosophers at home today
	philosopher := []string{"Socrates", "Locke", "Plato", "Krato", "Aristotle"}

	//	Specify the nunber of goroutines to wait on
	wg.Add(len(philosopher))
	
	// Each philosopher is going to the table
	// Loop through the philosopher slice and send each one to the table
	for i, v := range philosopher {

		// Anonymous function to create a functionality and control flow
		func(index int, value string) {
			var leftStick int
			var rightStick int



			if index == len(philosopher) - 1 {
				rightStick = len(philosopher)
				leftStick = len(philosopher) - index
			} else {
				rightStick = index + 1
				leftStick = index + 2
			}

			go philEat(value, rightStick, leftStick)

		} (i, v) 
	}

	wg.Wait()

}
	
func philEat (phil string, rS int, lS int) {

	// Whois at the table
	fmt.Printf("%s is at the table and ready to eat", phil)

	time.Sleep(time.Duration(period.wait) * time.Second)

	fmt.Println("stuff")
	wg.Done()
}
