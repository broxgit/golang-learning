package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepProcess() {
	fmt.Println("Start")
	go func() {
		fmt.Println("processing")
	}()
	time.Sleep(time.Millisecond * 10) // this is bad, DON'T DO THIS!
	fmt.Println("done")
}

func demoNoSynchronization() {
	var counter = 0
	for i := 0; i < 20; i++ {
		go func() {
			counter++
			fmt.Println(counter)
		}()
	}
	time.Sleep(time.Millisecond * 10)
}

// A mutex serializes access to the code under lock. The reason we simply define our lock as 'lock sync.Mutex' is because the default value of a sync.Mutex is unlocked
func demoSyncMutex() {
	var (
		counter = 0
		lock    sync.Mutex
	)

	for i := 0; i < 20; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			counter++
			fmt.Println(counter)
		}()
	}
	time.Sleep(time.Millisecond * 10)
}

func demoDeadlock() {
	var (
		lock sync.Mutex
	)

	go func() {
		lock.Lock()
	}()
	time.Sleep(time.Millisecond * 10)
	lock.Lock()
}

func main() {
	//sleepProcess()
	//demoNoSynchronization()
	//demoSyncMutex()
	demoDeadlock()
}
