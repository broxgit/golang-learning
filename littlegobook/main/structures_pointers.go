package main

import "fmt"

// Go equivalent of a class -- called a structure
type Saiyan struct {
	Name  string
	Power int
}

// We can associate a method with a structure, the structure will receive the method
func (s *Saiyan) Super() {
	s.Power += 10000
}

// A pointer is a memory address; it's the location of where to find the actual value
// Go passes arguments to functions as copies, unless you use a pointer then it passes the address of the argument/object
// Pre-pending a variable name with * tells Go that you want the pointer to value of the argument being passed through
func pointerOfValue(s *Saiyan) {
	s.Power += 10000
}

// Demonstrating how Go passes the argument as a copy, increasing the value of Power won't actually have an effect on the argument being passed to this function
func copyOfValue(s Saiyan) {
	s.Power += 10000
}

// since this function will receive a COPY of the address to the variable 's', no changes will actually be made to the argument value being passed through
func reassignPointerAddress(s *Saiyan) {
	s = &Saiyan{"Gohan", 1000}
}

func pointerDemonstration() {
	// sending a copy of goku to a function to see if the value of Power changes:
	goku := Saiyan{"Goku", 9000}
	copyOfValue(goku)
	fmt.Println("Sending a copy of our value to a function th at will attempt to change the value of Power (9000 to 19000):")
	fmt.Println(goku.Power)

	// Sending a pointer of gokuPointer to a function to see if the value of Power changes:
	// pre-pending a variable with & gets the address of our value
	gokuPointer := &Saiyan{"Goku", 9000}
	pointerOfValue(gokuPointer)
	fmt.Println("Sending a pointer value to a function that will change the value of Power using the pointer's address for the argument (9000 to 19000):")
	fmt.Println(gokuPointer.Power)

	// Sending a pointer value to a function that will attempt to reassign the address of the value:
	// When sending a pointer to a function, we are sending a copy of the address to the function, not the actual address
	gokuAddChange := &Saiyan{"Goku", 9000}
	reassignPointerAddress(gokuAddChange)
	fmt.Println("Sending a pointer value to a function that will attempt to reassign the address of the argument (9000 to 10000):")
	fmt.Println(gokuAddChange.Power)
}

func structureFunctionDemo() {
	goku := &Saiyan{"Goku", 9001}
	goku.Super()
	fmt.Println(goku.Power)
}

func main() {
	//pointerDemonstration()
	structureFunctionDemo()

}
