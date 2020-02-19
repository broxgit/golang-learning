package main

import "fmt"

func printStuff() {
	fmt.Println("I'm deferring stuff!")
}

func main() {
	defer printStuff()
	if true {
		fmt.Println("I should be before the deferment!")
		return
	}
}
