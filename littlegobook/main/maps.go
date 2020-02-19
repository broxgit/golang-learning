package main

import "fmt"

func demoMaps() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001

	// attempt finding a value in the map using the key "vegeta" and assigning the value to power.
	// exists will return true or false
	power, exists := lookup["vegeta"]

	// prints 0, false
	fmt.Println(power, exists)

	lookup["vegeta"] = 9002
	power, exists = lookup["vegeta"]
	// prints 9002, true
	fmt.Println(power, exists)

	delete(lookup, "vegeta")
	power, exists = lookup["vegeta"]
	// prints 0, false
	fmt.Println(power, exists)

	total := len(lookup)
	fmt.Printf("Total: %d", total)

	// maps grow dynamically, however we can still supply a second argument to 'make' to set an initial size:
	lookup = make(map[string]int, 100)
	// if you have an idea of how many keys your map will have, defining an initial size can help with performance

	lookup["bartsimpson"] = 10000
	lookup["generalpatton"] = 11000

	for key, value := range lookup {
		fmt.Printf("\nKey: %s -- Value: %d", key, value)
	}

}

func main() {
	demoMaps()
}
