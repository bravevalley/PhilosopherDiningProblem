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

// var stickA sync.Mutex
// var stickB sync.Mutex
// var stickC sync.Mutex
// var stickD sync.Mutex
// var stickE sync.Mutex

// var sticks = map[int]sync.Mutex {
// 	1: stickA,
// 	2: stickB,
// 	3: stickC,
// 	4: stickD,
// 	5: stickD,
// }

var wg sync.WaitGroup

func main() {
	rightStick := &sync.Mutex{}

	// Set the table - Hypothetical
	fmt.Println("The table is set!")

	// The philosophers at home today
	philosopher := []string{"Socrates", "Locke", "Plato", "Krato", "Aristotle"}

	//	Specify the nunber of goroutines to wait on
	wg.Add(len(philosopher))

	// Each philosopher is going to the table
	// Loop through the philosopher slice and send each one to the table
	for i:=0; i< len(philosopher); i++ {

		leftStick := &sync.Mutex{}

		go philEat(philosopher[i], rightStick, leftStick)

		rightStick = leftStick

	}

	wg.Wait()

}

func philEat(phil string, rS, lS *sync.Mutex) {

	// Whois at the table
	fmt.Printf("%s is at the table and ready to eat\n", phil)

	time.Sleep(time.Duration(period.wait) * time.Second)

	// Philosopher May begin eatting

	for i := 1; i <= feeding; i++ {

		// Lock the stick with the one eating

		rS.Lock()
		fmt.Printf("\t%s has picked up the right stick\n", phil)

		// Update the stick in use
		// Pick up the sticks to eat
		lS.Lock()
		fmt.Printf("\t%s has picked up the stick left\n", phil)

		// Sleep for a Second
		time.Sleep(1 * time.Second)

		// Time to eat
		fmt.Printf("%s has both sticks --- now eating...\n", phil)
		time.Sleep(time.Duration(period.eat) * time.Second)

		// Drop the sticks
		fmt.Printf("\t\t%s JUST DROPPED the right stick\n", phil)
		rS.Unlock()
		fmt.Printf("\t\t%s JUST DROPPED the left stick\n", phil)
		lS.Unlock()

	}

	// Time to think

	fmt.Println(phil, "is thinking..... ")
	time.Sleep(time.Duration(period.think) * time.Second)

	fmt.Printf("\t\t%s has finished eating...\n", phil)
	wg.Done()
}
