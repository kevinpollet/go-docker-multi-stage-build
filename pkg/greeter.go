package pkg

import "fmt"

// Greet greets the give name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}