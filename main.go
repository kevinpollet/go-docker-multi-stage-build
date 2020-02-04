package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("id:", uuid.New())
	fmt.Println("message:", "Hello, World!")
}
