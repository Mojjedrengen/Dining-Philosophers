package main

import "fmt"

func forkLogic(messege string, isTaken bool) {

}

var ch1 = make(chan string)
var ch2 = make(chan string)

func fork1() {

}

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
		}
	}
}

func main() {
	fmt.Println("Hello World")
}
