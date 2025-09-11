package main

import "fmt"

type PhilosoferState int

const (
	Eating PhilosoferState = iota
	Thinking
)

var PhilosoferStateName = map[PhilosoferState]string{
	Eating:   "Eating",
	Thinking: "Thinking",
}

var ch1 = make(chan string)
var ch2 = make(chan string)

func forkgeneric(resive <-chan string, leftmsg chan<- string, rightmsg chan<- string) {
	var isTaken bool = false

	for true {

		var msg = <-resive

		fmt.Printf("fork got msg: %s and isTaken == %t\n", msg, isTaken)

		if msg == "left" {
			if isTaken == false {
				isTaken = true
				leftmsg <- "grabed"
			} else {
				leftmsg <- "is taken"
			}
		} else if msg == "right" {
			if isTaken == false {
				isTaken = true
				rightmsg <- "grabed"
			} else {
				rightmsg <- "is taken"
			}
		} else if msg == "leave" {
			isTaken = false
		}
	}
}

func philosofGeneric(leftFork chan<- string, rightFork chan<- string, leftForkMsg <-chan string, rightForkMsg <-chan string, philosofer int) {
	var timesEaten int = 0
	var state PhilosoferState = Thinking
	var haveLeftFork bool = false
	var haveRightFork bool = false

	for true {
		if state == Thinking {
			leftFork <- "left"
			var leftmsg = <-leftForkMsg
			rightFork <- "right"
			var rightmsg = <-rightForkMsg

			if leftmsg == "grabed" {
				haveLeftFork = true
			}
			if rightmsg == "grabed" {
				haveRightFork = true
			}

			fmt.Printf("%d: rightFork == %t, leftFork == %t\n", philosofer, haveLeftFork, haveRightFork)

			if haveLeftFork == true && haveRightFork == true {
				state = Eating
				timesEaten++
				fmt.Printf("%d is now %s\n", philosofer, PhilosoferStateName[state])
				fmt.Printf("%d has eaten %d times\n", philosofer, timesEaten)
			} else if haveLeftFork == true && haveRightFork == false {
				leftFork <- "leave"
			} else if haveLeftFork == false && haveRightFork == true {
				rightFork <- "leave"
			}

		} else if state == Eating {
			leftFork <- "leave"
			rightFork <- "leave"
			state = Thinking
			fmt.Printf("%d is now %s\n", philosofer, PhilosoferStateName[state])
		}
	}
}

var fork1 = make(chan string)
var fork2 = make(chan string)
var fork3 = make(chan string)
var fork4 = make(chan string)
var fork5 = make(chan string)

var philosof1L = make(chan string)
var philosof2L = make(chan string)
var philosof3L = make(chan string)
var philosof4L = make(chan string)
var philosof5L = make(chan string)
var philosof1R = make(chan string)
var philosof2R = make(chan string)
var philosof3R = make(chan string)
var philosof4R = make(chan string)
var philosof5R = make(chan string)

func main() {
	fmt.Println("Hello World")

	go forkgeneric(fork1, philosof1L, philosof5R)
	go forkgeneric(fork2, philosof2L, philosof1R)
	go forkgeneric(fork3, philosof3L, philosof2R)
	go forkgeneric(fork4, philosof4L, philosof3R)
	go forkgeneric(fork5, philosof5L, philosof4R)

	go philosofGeneric(fork2, fork1, philosof1L, philosof1R, 1)
	go philosofGeneric(fork3, fork2, philosof2L, philosof2R, 2)
	go philosofGeneric(fork4, fork3, philosof3L, philosof3R, 3)
	go philosofGeneric(fork5, fork4, philosof4L, philosof4R, 4)
	go philosofGeneric(fork1, fork5, philosof5L, philosof5L, 5)

	for true {
	}

}
