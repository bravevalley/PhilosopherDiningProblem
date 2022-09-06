package main

import (
	"fmt"
	"sync"
	"time"
)

// Number of times to eat the meal
const feeding = 3

var period = struct {
	think int
	eat   int
	wait  int
}{
	think: 2,
	eat:   3,
	wait:  2,
}

var sticks = map[int]string {
	1: "free",
	2: "free",
	3: "free",
	4: "free",
	5: "free",
}

var wg sync.WaitGroup
var mu sync.Mutex

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

			if index == len(philosopher)-1 {
				rightStick = len(philosopher)
				leftStick = len(philosopher) - index
			} else {
				rightStick = index + 1
				leftStick = index + 2
			}

			go philEat(value, rightStick, leftStick)

		}(i, v)
	}

	wg.Wait()

}

func philEat(phil string, rS int, lS int) {

	// Whois at the table
	fmt.Printf("%s is at the table and ready to eat\n", phil)

	time.Sleep(time.Duration(period.wait) * time.Second)

	// Philosopher May begin eatting

	for i := 1; i < feeding; i++ {

		// Lock the stick with the one eating
		mu.Lock()

		// Pick up the sticks to eat
		fmt.Printf("\t%s has picked up stick %d and %d\n", phil, rS, lS)

		// Sleep for a Second
		time.Sleep(1 * time.Second)

		// Update the stick in use
		sticks[rS] = fmt.Sprintln("In use")
		sticks[lS] = fmt.Sprintln("In use")


		// Time to eat
		time.Sleep(time.Duration(period.eat) * time.Second)

		

		// Drop the sticks
		sticks[rS] = fmt.Sprintln("Free")
		sticks[lS] = fmt.Sprintln("Free")

		fmt.Printf("\t%s JUST DROPPED stick %d and %d\n", phil, rS, lS)

		mu.Unlock()
		}

		// Time to think

		fmt.Println(phil, "is thinking..... ")
		time.Sleep(time.Duration(period.think) * time.Second)

		fmt.Printf("\t\t%s has finished eating...\n", phil)
	wg.Done()
}
