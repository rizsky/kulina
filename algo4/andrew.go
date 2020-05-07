package main

import "fmt"

/*
Andrew walks through 100 switches from point A to point B with 1 to 100 as the number.
At the first trip, Andrew presses all of the switches so lamps are on. Second trip, Andrew
only presses switches that multiplying of two. The next trip, Andrew presses switches
that multiplying of three. Andrew repeats his trips from A to B for 100 times. Write down
the code to determine how many lamps will turn on after 100 trips from A to B.
*/

func main() {
	trips(100)
}

func trips(x int) {
	lamps := 0

	for i := 1; i <= x; i++ {

		if i%3 == 0 {
			lamps += (x / 3)
			continue
		}

		if i%2 == 0 {
			lamps += (x / 2)
			continue
		}

		lamps += x
	}

	fmt.Printf("\rLamps that turn on : %d\n", lamps)
}
