package main

import (
	"fmt"
	"time"
)

/*
To make sure that the program does not lead to a deadlock we have 3 statements
(1) Even numbered Philosofers are left handed
(2) Odd numbered Philosofers are right handed
and (3) Philosofers are stubborn

Number (1) and (2) makes it so that the Philosofers alternates between which fork they want first
Because they alternate between the side they want to grab first, they free up Philosofer 5.

Number (3) means that once a Philosofer gets their main fork they would not release it before they get the other one.
In tandem with (1) and (2) this makes it so that the philosofers all gets to eat.

This does not make sure that the Philosofers eats an equal amount of food by it self.

To give the Philosofers time to think we make them wait for 0.5s before attempting to grab their main fork after they have put them down
Inadvertently this made it so that the Philosofers eat an equal amount of food.
To make it so that they do not wait while thinking comment out line 104
*/

type Calls int

const (
	Left Calls = iota
	Right
	Taken
	Free
	Leave
)

type PState int

const (
	Thinking PState = iota
	Eating
)

func forks(leftChan chan Calls, rightChan chan Calls, nr int) {
	var isTaken bool = false
	var msg Calls

	for true {
		select {
		case msg = <-leftChan:
		case msg = <-rightChan:
		}
		if isTaken && msg == Leave {
			isTaken = false
		} else if !isTaken && msg == Left {
			respond(leftChan, Free)
			isTaken = true
		} else if !isTaken && msg == Right {
			respond(rightChan, Free)
			isTaken = true
		} else if isTaken && msg == Left {
			respond(leftChan, Taken)
		} else if isTaken && msg == Right {
			respond(rightChan, Taken)
		}

	}
}
func respond(c chan<- Calls, msg Calls) {
	c <- msg
}

func philosof(domHandChan chan Calls, subHandChan chan Calls, nr int) {
	var timesEaten int = 0
	var domHand Calls
	var subHand Calls
	var state PState = PState(Thinking)
	if nr%2 != 0 {
		domHand = Left
		subHand = Right
	} else {
		domHand = Right
		subHand = Left
	}

	for true {
		if state == Thinking {
			respond(domHandChan, domHand)
			if <-domHandChan == Free {
				for state == Thinking {
					respond(subHandChan, subHand)
					if <-subHandChan == Free {
						state = Eating
						timesEaten += 1
						fmt.Printf("%d is now eating. They have eaten %d times.\n", nr, timesEaten)
					}
				}
			}
		} else {
			respond(domHandChan, Leave)
			respond(subHandChan, Leave)
			state = Thinking
			fmt.Printf("%d is now Thinking.\n", nr)
			time.Sleep(500 * time.Millisecond) // <- Comment out to make the Philosofers not wait while thinking
		}
	}
}

var cf11 = make(chan Calls)
var cf12 = make(chan Calls)
var cf21 = make(chan Calls)
var cf22 = make(chan Calls)
var cf31 = make(chan Calls)
var cf32 = make(chan Calls)
var cf41 = make(chan Calls)
var cf42 = make(chan Calls)
var cf51 = make(chan Calls)
var cf52 = make(chan Calls)

func main() {
	go forks(cf11, cf12, 1)
	go forks(cf21, cf22, 2)
	go forks(cf31, cf32, 3)
	go forks(cf41, cf42, 4)
	go forks(cf51, cf52, 5)

	go philosof(cf11, cf52, 1)
	go philosof(cf12, cf21, 2)
	go philosof(cf31, cf22, 3)
	go philosof(cf32, cf41, 4)
	go philosof(cf51, cf42, 5)

	for true {

	}
}
