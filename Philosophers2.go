package main

import (
	"fmt"
	"time"
)

//lige er venstre håndet
//ulige er højre
//philosofer er stædige

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

func forks(leftChan chan Calls, rightChan chan Calls) {
	var isTaken bool = false
	var msg Calls

	for true {
		select {
		case msg = <-leftChan:
		case msg = <-rightChan:
		}
		fmt.Printf("%d and is taken = %b\n", msg, isTaken)
		if isTaken && msg == Leave {
			isTaken = false
		} else if !isTaken && msg == Left {
			respond(leftChan, Free)
			isTaken = true
			fmt.Printf("sent msg free to left\n")
		} else if !isTaken && msg == Right {
			respond(rightChan, Free)
			isTaken = true
			fmt.Printf("sent msg free to right\n")
		} else if isTaken && msg == Left {
			respond(leftChan, Taken)
			fmt.Printf("sent msg taken to left\n")
		} else if isTaken && msg == Right {
			respond(rightChan, Taken)
			fmt.Printf("sent msg taken to left\n")
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
	if nr%2 == 0 {
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
				fmt.Printf("%d got main fork", nr)
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
			fmt.Printf("%d is now Thinking.\n", nr)
			time.Sleep(500 * time.Millisecond)
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
	go forks(cf11, cf12)
	go forks(cf21, cf22)
	go forks(cf31, cf32)
	go forks(cf41, cf42)
	go forks(cf51, cf52)

	go philosof(cf11, cf52, 1)
	go philosof(cf12, cf21, 2)
	go philosof(cf31, cf22, 3)
	go philosof(cf32, cf41, 4)
	go philosof(cf51, cf42, 5)

	for true {

	}
}
