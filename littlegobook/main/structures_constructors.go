package main

import "fmt"

type Bandsaw struct {
	Brand      string
	HorsePower int
}

// Structures don’t have constructors. Instead, you create a function that returns an instance of the desired type (like a factory):
func NewBandsaw(brand string, horsepower int) *Bandsaw {
	return &Bandsaw{
		Brand:      brand,
		HorsePower: horsepower,
	}
}

// Despite the lack of constructors, Go does have a built-in new function which is used to allocate the memory required
// by a type. The result of new(X) is the same as &X{}:
func demonstrateKeywordNEW() {
	grizzly17 := new(Bandsaw)
	grizzly17.Brand = "Grizzly"
	grizzly17.HorsePower = 2
	// same as:
	grizzly19 := &Bandsaw{
		Brand:      "Grizzly",
		HorsePower: 2,
	}

	// Which you use is up to you, but you’ll find that most people prefer the latter whenever they have fields to initialize,
	// since it tends to be easier to read

	fmt.Println(grizzly17)
	fmt.Println(grizzly19)
}
