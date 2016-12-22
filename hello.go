package main

import (
	"fmt"
)

// Greeting greets someone.
func Greeting(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Greeting("Caleb"))
	fmt.Println("So cool!")
}
