package main

import "fmt"

func demoArrays() {
	scores := [4]int{1, 2, 3, 4}

	for index, value := range scores {
		fmt.Print(index)
		fmt.Print(value)
	}
}

// In Go, you rarely, if ever, use arrays directly. Instead, you use slices.
// A slice is a lightweight structure that wraps and represents a portion of an array
func demoSlices() {
	// This is a slight variation to the way we declared the array in the above method, however slices don't require a length within the square brackets
	scores := []int{1, 2, 3, 4}
	for index, value := range scores {
		fmt.Print(index)
		fmt.Print(value)
	}

	fmt.Println()

	// We use 'make' instead of 'new' because there's more to creating a slice than just allocating memory (which is what 'new' does)
	// Specifically we have to allocate the memory to the underlying array and also initialize the slice
	// Here we will initialize a slice with a length of 10 and a capacity of 10
	scoresMakeOne := make([]int, 10)
	scoresMakeOne[7] = 7
	// the following will print [0 0 0 0 0 0 0 7 0 0]
	fmt.Println(scoresMakeOne)

	// Here we are creating a slice with a length of 0 but with a capacity of 10
	scoresMakeTwo := make([]int, 0, 10)
	scoresMakeTwo = append(scoresMakeTwo, 7)
	// the following will print [7]
	fmt.Println(scoresMakeTwo)

	// we can't assign the index 7 with a value of 7 like we did to scoresMakeOne because the length is currently 0 so the following would break the code:
	// code: scoresMakeTwo[7] = 7
	// if we wanted to accomplish the same thing as we did with scoresMakeOne, we could do the following:
	scoresMakeTwo = scoresMakeTwo[0:8]
	scoresMakeTwo[7] = 7

	// the following will print [7 0 0 0 0 0 0 7]
	fmt.Println(scoresMakeTwo)
}

// In Go, if the array gets full, it will create a new larger array and copy the values over
// Go grows arrays with a 2x algorithm, meaning if we have a capacity of 10 but we are trying to assign a value to index 11, the capacity will grow to 20
func demoSliceGrowth() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		// if our capacity has changed, Go will grow our array to accommodate the new data
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}

func demoSliceInitializers() {
	// In Go, there are ultimately four common ways to initialize a slice:

	// Use this method when the values are known ahead of time:
	names := []string{"brock", "patton", "erin", "eden"}

	// Use this method when you'll be writing to specific indexes of a slice:
	checks := make([]bool, 10)

	// This is a nil slice and is used in conjunction with append, when the number of elements is unknown:
	var tools []string

	// This version lets you specify a capacity, if you have a general idea of how many elements you'll need:
	scores := make([]int, 0, 20)

	fmt.Printf("%d %d %d %d", names, checks, tools, scores)
}

func demoSlicingASlice() {
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	slice[0] = 999

	// The following will output [1 2 999 4 5], the value of scores HAS changed
	fmt.Println(scores)
}

func demoCopyingASlice() {
	scores := []int{1, 2, 3, 4, 5}
	slice := make([]int, 5)
	copy(slice, scores[2:4])
	slice[0] = 999

	// The following will output [1 2 3 4 5], the value of scores hasn't changed
	fmt.Println(scores)
}

func main() {
	//demoSlices()
	//demoSliceGrowth()
	//demoSliceInitializers()
	//demoSlicingASlice()
	demoCopyingASlice()
}
