package main

import (
	"fmt"
	"time"
)

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

func forkgeneric(resive <-chan string, leftmsg chan<- string, rightmsg chan<- string, fork int) {
	var isTaken bool = false

	for true {

		var msg = <-resive
		fmt.Printf("f%d got msg: %s\n", fork, msg)
		//		fmt.Printf("fork got msg: %s and isTaken == %t\n", msg, isTaken)

		if msg == "left" {
			if isTaken == false {
				isTaken = true
				leftmsg <- "grabed"
				fmt.Printf("f%d: graped left\n", fork)
			} else {
				leftmsg <- "is taken"
				fmt.Printf("f%d: denied left\n", fork)
			}
		} else if msg == "right" {
			if isTaken == false {
				isTaken = true
				rightmsg <- "grabed"
				fmt.Printf("f%d: graped right\n", fork)
			} else {
				rightmsg <- "is taken"
				fmt.Printf("f%d: denied right\n", fork)
			}
		} else if msg == "leave" {
			isTaken = false
		}
	}
}

/*
To fix deadlock Philosofers have these 3 traits:
lige er venstre håndet
ulige er højre
philosofer er stædige
*/
func philosofGeneric(leftFork chan<- string, rightFork chan<- string, leftForkMsg <-chan string, rightForkMsg <-chan string, philosofer int, leftForkRank int, rightForkRank int) {
	var timesEaten int = 0
	var state PhilosoferState = Thinking
	var haveLeftFork bool = false
	var haveRightFork bool = false

	for true {
		if state == Thinking {
			var leftmsg string
			var rightmsg string

			if philosofer%2 == 0 {
				fmt.Printf("%d is attempting go get main: left: %d\n", philosofer, leftForkRank)
				leftFork <- "left"
				fmt.Printf("%d sent msg\n", philosofer)
				leftmsg = <-leftForkMsg
				fmt.Printf("%d attempted to gain main. got msg: %s\n", philosofer, leftmsg)
				if leftmsg == "grabed" {
					haveLeftFork = true
					for !haveRightFork {
						rightFork <- "right"
						rightmsg = <-rightForkMsg
						fmt.Printf("%d is attempting to gain second fork\n", philosofer)
						if rightmsg == "grabed" {
							haveRightFork = true
						}
					}
				}
			} else {
				fmt.Printf("%d is attempting go get main: right: %d\n", philosofer, rightForkRank)
				rightFork <- "right"
				fmt.Printf("%d sent msg\n", philosofer)
				rightmsg = <-rightForkMsg
				fmt.Printf("%d attempted to gain main. got msg: %s\n", philosofer, rightmsg)
				if rightmsg == "grabed" {
					haveRightFork = true
					for !haveLeftFork {
						leftFork <- "left"
						leftmsg = <-leftForkMsg
						fmt.Printf("%d is attempting to gain second fork\n", philosofer)
						if leftmsg == "grabed" {
							haveLeftFork = true
						}
					}
				}
			}
			/*
				fmt.Printf("%d attepmts to send\n", philosofer)
				leftFork <- "left"
				fmt.Printf("%d sent message to left: %d\n", philosofer, leftForkRank)
				leftmsg = <-leftForkMsg
				fmt.Printf("%d: got message %s from left: %d\n", philosofer, leftmsg, leftForkRank)
				rightFork <- "right"
				fmt.Printf("%d sent message to right: %d\n", philosofer, rightForkRank)
				rightmsg = <-rightForkMsg
			*/
			/*
				if leftForkRank < rightForkRank {
					fmt.Printf("%d attepmts to send\n", philosofer)
					leftFork <- "left"
					fmt.Printf("%d sent message to left: %d\n", philosofer, leftForkRank)
					leftmsg = <-leftForkMsg
					fmt.Printf("%d: got message %s from left: %d\n", philosofer, leftmsg, leftForkRank)
					rightFork <- "right"
					fmt.Printf("%d sent message to right: %d\n", philosofer, rightForkRank)
					rightmsg = <-rightForkMsg
					fmt.Printf("%d: got message %s from right: %d\n", philosofer, rightmsg, rightForkRank)
				} else if leftForkRank > rightForkRank {
					fmt.Printf("%d attepmts to send\n", philosofer)
					rightFork <- "right"
					fmt.Printf("%d sent message to right: %d\n", philosofer, rightForkRank)
					rightmsg = <-rightForkMsg
					fmt.Printf("%d: got message %s from right: %d\n", philosofer, rightmsg, rightForkRank)
					leftFork <- "left"
					fmt.Printf("%d sent message to left: %d\n", philosofer, leftForkRank)
					leftmsg = <-leftForkMsg
					fmt.Printf("%d: got message %s from left: %d\n", philosofer, leftmsg, leftForkRank)
				} else {
					fmt.Printf("%d: rank invalid\n", philosofer)
					break
				}
			*/
			//fmt.Printf("%d got message %s from left and %s from right\n", philosofer, leftmsg, rightmsg)
			fmt.Printf("%d: leftFork == %d is %t, rightFork == %d is %t\n", philosofer, leftForkRank, haveLeftFork, rightForkRank, haveRightFork)

			if haveLeftFork == true && haveRightFork == true {
				state = Eating
				timesEaten++
				fmt.Printf("%d is now %s\n", philosofer, PhilosoferStateName[state])
				fmt.Printf("%d has eaten %d times\n", philosofer, timesEaten)
			} /* else if haveLeftFork == true && haveRightFork == false {
				leftFork <- "leave"
				fmt.Printf("%d put left fork: %d down\n", philosofer, leftForkRank)
			} else if haveLeftFork == false && haveRightFork == true {
				rightFork <- "leave"
				fmt.Printf("%d put right fork: %d down\n", philosofer, rightForkRank)
			}*/

		} else if state == Eating {
			leftFork <- "leave"
			haveLeftFork = false
			rightFork <- "leave"
			haveRightFork = false
			state = Thinking
			fmt.Printf("%d is now %s\n", philosofer, PhilosoferStateName[state])
		}
		fmt.Printf("%d now sleeps\n", philosofer)
		time.Sleep(1000 * time.Millisecond)
	}
}

var fork1 = make(chan string, 10)
var fork2 = make(chan string, 10)
var fork3 = make(chan string, 10)
var fork4 = make(chan string, 10)
var fork5 = make(chan string, 10)

var philosof1L = make(chan string, 10)
var philosof2L = make(chan string, 10)
var philosof3L = make(chan string, 10)
var philosof4L = make(chan string, 10)
var philosof5L = make(chan string, 10)
var philosof1R = make(chan string, 10)
var philosof2R = make(chan string, 10)
var philosof3R = make(chan string, 10)
var philosof4R = make(chan string, 10)
var philosof5R = make(chan string, 10)

func main() {
	fmt.Println("Hello World")

	go forkgeneric(fork1, philosof1L, philosof5R, 1)
	go forkgeneric(fork2, philosof2L, philosof1R, 2)
	go forkgeneric(fork3, philosof3L, philosof2R, 3)
	go forkgeneric(fork4, philosof4L, philosof3R, 4)
	go forkgeneric(fork5, philosof5L, philosof4R, 5)

	go philosofGeneric(fork2, fork1, philosof1L, philosof1R, 1, 2, 1)
	go philosofGeneric(fork3, fork2, philosof2L, philosof2R, 2, 3, 2)
	go philosofGeneric(fork4, fork3, philosof3L, philosof3R, 3, 4, 3)
	go philosofGeneric(fork5, fork4, philosof4L, philosof4R, 4, 5, 4)
	go philosofGeneric(fork1, fork5, philosof5L, philosof5R, 5, 1, 5)

	var debug = 0
	for true {
		debug++
		if debug == 2500000000 {
			//	break
		}
	}

}
