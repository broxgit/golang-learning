package main

import "fmt"

type Logger interface {
	Log(message string)
}

// Implementing an interface in Go is implicit. If your structure has a function name Log with a string parameter and no return value,
// then it can be used as a Logger. This cuts down on the verboseness of using interfaces:
type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
