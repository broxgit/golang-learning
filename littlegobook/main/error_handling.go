package main

import (
	"errors"
	"fmt"
)

func strConv(num int) error {
	if num > 10 {
		fmt.Println("not a valid number")
		return errors.New("Invalid number")
	} else {
		fmt.Println(num)
		return nil
	}
}

func main() {
	fmt.Println(strConv(999))
}
