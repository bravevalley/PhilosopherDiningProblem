package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var period = struct {
	// Number of times to eat the meal
	feeding int
	// Time it takes to think
	think int
	// Time it takes to eat
	eat int
	// Time to wait
	wait int
}{
	feeding: 3,
	think:   2,
	eat:     3,
	wait:    2,
}

// Initialize a WaitGroup
var wg sync.WaitGroup

// Keeping track of order
var position []string

// Lock the slice
var mu sync.Mutex

func main() {

	// Create the right Stick as a pointer type to sync.Mutex interface so we can call the lock function on the variable.
	rightStick := &sync.Mutex{}

	// Set the table - Hypothetical
	fmt.Println("The table is set!")

	// The philosophers at home today
	philosopher := []string{"Socrates", "Locke", "Plato", "Krato", "Aristotle"}

	//	Specify the nunber of goroutines to wait on
	wg.Add(len(philosopher))

	// Each philosopher is going to the table
	// Loop through the philosopher slice and send each one to the table
	for i := 0; i < len(philosopher); i++ {

		// Create the let Stick as a pointer type to sync.Mutex interface so we can call the lock function on the variable.
		leftStick := &sync.Mutex{}

		// Create a goroutine in which all the philosophers will try to call the lock method of each stick
		// If a philosopher is able to call lock on the the first stick then the routine stops until
		// it is able to call the lock on the second stick
		go philEat(philosopher[i], rightStick, leftStick)

		// Make right stick the left to stimulate rotation
		rightStick = leftStick

	}

	wg.Wait()

	for i, v := range position {
		// Convert index to string
		pos := strconv.Itoa(i + 1)
		var fin string

		// Test string suffix to append the positonal order
		if strings.HasSuffix(pos, "1") {
			fin = fmt.Sprintf("%sst", pos)
		} else if strings.HasSuffix(pos, "2") {
			fin = fmt.Sprintf("%snd", pos)
		} else if strings.HasSuffix(pos, "3") {
			fin = fmt.Sprintf("%srd", pos)
		} else {
			fin = fmt.Sprintf("%sth", pos)
		}

		fmt.Printf("%s finished %s\n", v, fin)
	}

}

func philEat(phil string, rS, lS *sync.Mutex) {

	// Whois at the table
	fmt.Printf("%s is at the table and ready to eat\n", phil)

	time.Sleep(time.Duration(period.wait) * time.Second)

	// Philosopher May begin eatting

	for i := 1; i <= period.feeding; i++ {

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

	mu.Lock()
	position = append(position, phil)
	mu.Unlock()

	wg.Done()
}
