package main

import "fmt"

type Tool struct {
	Name string
}

func (t *Tool) Introduce() {
	fmt.Printf("Hi, I'm a %s\n", t.Name)
}

type TableSaw struct {
	*Tool
	Model string
}

func main() {
	saw := &TableSaw{
		Tool:  &Tool{"Table saw"},
		Model: "G0917",
	}

	// We now have access to the introduce method of Tool:
	saw.Introduce()

	// Or we can get the value of Name directly like this:
	fmt.Println(saw.Name)
	// Or
	fmt.Println(saw.Tool.Name)

	// Both of the above are perfectly valid
}
