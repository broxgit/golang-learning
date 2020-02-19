package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}

func demoChannels() {
	c := make(chan int, 100)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	for {
		c <- rand.Int()
		fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}
}

func demoBufferedChannels() {
	c := make(chan int, 100)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
}

// Even with buffering, there comes a point where we need to start dropping messages. We can’t use up an infinite
// amount of memory hoping a worker frees up. For this, we use Go’s select.
// Syntactically, select looks a bit like a switch. With it, we can provide code for when the channel isn’t available to send to.
func demoChannelSelect() {
	c := make(chan int, 100)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	for {
		select {
		case c <- rand.Int():
			// optional code here
		default:
			// this can be left empty to silently drop the data
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

// Another popular option is to timeout. We’re
// willing to block for some time, but not forever. This is also something easy to achieve in Go.
func demoChannelTimeout() {
	// To block for a maximum amount of time, we can use the time.After function.
	c := make(chan int, 100)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	for {
		select {
		case c <- rand.Int():
			// optional code here
		case <-time.After(time.Millisecond * 100):
			fmt.Println("timed out")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	//demoChannels()
	//demoBufferedChannels()
	demoChannelSelect()
}
