package main

import "fmt"

func main() {
	stra := "good boy is a good boy"
	fmt.Printf("\nstra: %s -- Length: %d\n", stra, len(stra))
	for i := 0; i < len(stra); i++ {
		fmt.Printf("%d  ", stra[i])
	}
	byts := []byte(stra)
	fmt.Printf("\nBytes: %s\n", byts)
	for i := 0; i < len(byts); i++ {
		fmt.Printf("%d  ", byts[i])
	}
	strb := string(byts)
	fmt.Printf("\nstrb: %s", strb)
}
