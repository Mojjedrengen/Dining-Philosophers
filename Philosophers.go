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

			if haveLeftFork == true && haveRightFork == true {
				state = Eating
				timesEaten++
				fmt.Println("%d is now %s", philosofer, PhilosoferStateName[state])
				fmt.Println("%d has eaten %d times", philosofer, timesEaten)
			} else if haveLeftFork == true && haveRightFork == false {
				leftFork <- "leave"
			} else if haveLeftFork == false && haveRightFork == true {
				rightFork <- "leave"
			}

		} else if state == Eating {
			leftFork <- "leave"
			rightFork <- "leave"
			state = Thinking
			fmt.Println("%d is now %s", philosofer, PhilosoferStateName[state])
		}
	}
}

var fork1 = make(chan string)
var fork2 = make(chan string)
var fork3 = make(chan string)
var fork4 = make(chan string)
var fork5 = make(chan string)

var philosof1 = make(chan string)
var philosof2 = make(chan string)
var philosof3 = make(chan string)
var philosof4 = make(chan string)
var philosof5 = make(chan string)

func main() {
	fmt.Println("Hello World")

	go forkgeneric(fork1, philosof1, philosof5)
	go forkgeneric(fork2, philosof2, philosof1)
	go forkgeneric(fork3, philosof3, philosof2)
	go forkgeneric(fork4, philosof4, philosof3)
	go forkgeneric(fork5, philosof5, philosof4)

	go philosofGeneric()

}
