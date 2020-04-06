package testmod

import "fmt"

// Hi - Say Hi <name>
func Hi(name string) {
	fmt.Printf("Hi %s\n", name)
}

// Hi2 - Say Hi <name>, <lastName>
func Hi2(name string, lastName string) {
	fmt.Printf("Hello %s, %s\n", name, lastName)
}
